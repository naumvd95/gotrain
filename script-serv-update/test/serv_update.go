package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"strings"
	"sync"

	"gopkg.in/yaml.v2"
)

var inputData = `
- name: dc1
  racks:
    - name: rack_1
      online: true
      servers:
        - name: s1
          sw_version: 1
        - name: s2
          sw_version: 1
        - name: s3
          sw_version: 1
        - name: s4
          sw_version: 1
        - name: s5
          sw_version: 1
    - name: rack_2
      online: true
      servers:
        - name: s1
          sw_version: 1
        - name: s2
          sw_version: 1
        - name: s3
          sw_version: 1
        - name: s4
          sw_version: 1
        - name: s5
          sw_version: 1
- name: dc2
  racks:
    - name: rack_1
      online: true
      servers:
        - name: s1
          sw_version: 1
        - name: s2
          sw_version: 1
        - name: s3
          sw_version: 1
        - name: s4
          sw_version: 1
        - name: s5
          sw_version: 555
    - name: rack_2
      online: true
      servers:
        - name: s1
          sw_version: 1
        - name: s2
          sw_version: 1
        - name: s3
          sw_version: 1
        - name: s4
          sw_version: 666
        - name: s5
          sw_version: 555
`

type Datacenter struct {
	Name  string `yaml:"name"`
	Racks []Rack `yaml:"racks"`
}

type Rack struct {
	Name    string   `yaml:"name"`
	Online  bool     `yaml:"online"`
	Servers []Server `yaml:"servers"`
}

type Server struct {
	Name    string `yaml:"name"`
	Version int    `yaml:"sw_version"`
}

func updateRack(rack Rack, newVersion int) map[string]error {
	servers := rack.Servers
	errorMap := make(map[string]error)

	// prepare routine boilerplate
	var wg sync.WaitGroup
	wg.Add(len(servers))

	for i := 0; i < len(servers); i++ {
		go func(s *Server) {
			defer wg.Done()
			servMeta := fmt.Sprintf("Server %v rack %v, updating from %v to %v", s.Name, rack.Name, s.Version, newVersion)
			fmt.Println(servMeta)

			oldVersion := s.Version
			if oldVersion > newVersion {
				errMsg := fmt.Sprintf("[failed]: %v: version mismatch\n", servMeta)
				servHash := fmt.Sprintf("%v_%v", s.Name, rack.Name)

				errorMap[servHash] = errors.New(errMsg)
			} else {
				s.Version = newVersion
			}
		}(&servers[i])
	}

	wg.Wait()

	return errorMap
}

func main() {
	newVersion := 2
	s := strings.NewReader(inputData)
	dataBytes, err := ioutil.ReadAll(s)
	if err != nil {
		panic(err)
	}

	var hardware []Datacenter
	err = yaml.Unmarshal(dataBytes, &hardware)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Current state of HW: %v\n", hardware)

	serverErrors := []map[string]error{}
	// dc loop can be multi-thread
	for _, dc := range hardware {
		for _, r := range dc.Racks {
			r.Online = false
			fmt.Printf("Updating rack %v online: %v in dc %v\n", r.Name, r.Online, dc.Name)

			errorsMap := updateRack(r, newVersion)
			if len(errorsMap) > 0 {
				fmt.Printf("Updating rack %v online: %v in dc %v: [failed], rack will be offline\n", r.Name, r.Online, dc.Name)
				serverErrors = append(serverErrors, errorsMap)
			} else {
				r.Online = true
			}
		}
	}
	if len(serverErrors) > 0 {
		fmt.Printf("Update failed: %v \n Current state of HW: %v\n", serverErrors, hardware)
		panic("Update failed")
	}

	fmt.Printf("[success] Current state of HW: %v\n", hardware)
}
