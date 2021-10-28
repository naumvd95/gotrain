package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"sync"
	"time"

	"gopkg.in/yaml.v2"
)

/* algo:
   1. read all file, otherwise we cant unmarshal it
   2. prepare data structs
   3. go routines while looping over DC's
   4. inside each routine looping over racks w/o concurrency
   5. mark rack offline
   6. inside each rack you may spin up go routine as well
   7. mark rack online
*/

type Server struct {
	Name    string `yaml:"name"`
	Version int    `yaml:"sw_version"`
}

type Rack struct {
	Name    string   `yaml:"name"`
	Status  string   `yaml:"status"`
	Servers []Server `yaml:"servers"`
}

type Datacenter struct {
	Name  string `yaml:"name"`
	Racks []Rack `yaml:"racks"`
}

func main() {
	inputFileName := "servers_input.yaml"
	data, err := ioutil.ReadFile(inputFileName)
	if err != nil {
		panic(err)
	}

	var dcs []Datacenter
	err = yaml.Unmarshal(data, &dcs)
	if err != nil {
		panic(err)
	}
	fmt.Printf("here is parsed data %v\n", dcs)

	var wg sync.WaitGroup
	wg.Add(len(dcs))
	for i := 0; i < len(dcs); i++ {
		currentDC := dcs[i]

		go func(d Datacenter) {
			fmt.Printf("We are in dc %v\n", d.Name)
			seconds := getRandSec(8)
			fmt.Printf("here is sec amount: %v\n", seconds)
			time.Sleep(time.Duration(seconds) * time.Second)
			fmt.Printf("first rack is %v\n", d.Racks[0].Name)

			defer wg.Done()
		}(currentDC)
	}
	wg.Wait()
}

func getRandSec(r int64) int64 {
	rand.Seed(time.Now().UnixNano())
	return rand.Int63n(r + 1)
}
