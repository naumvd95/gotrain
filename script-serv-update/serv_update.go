package main

import (
	"fmt"
	"io/ioutil"
	"sync"

	"gopkg.in/yaml.v2"
)

/*
logic:

ideal:
1. go through each DC simultaneously via go-routine
---
for now
2. in each DC go for-loop by racks -> set status to offline
3. in each rack spinup go-routine on all servers and add wait counter
4. after waiting all goroutines, set status to online and continue looping

*/

type Datacenter struct {
	Name  string `yaml:"name"`
	Racks []Rack `yaml:"racks"`
}

type Rack struct {
	Name    string   `yaml:"name"`
	Status  string   `yaml:"status"`
	Servers []Server `yaml:"servers"`
}

type Server struct {
	Name            string `yaml:"name"`
	SoftwareVersion int    `yaml:"sw_version"`
}

func getServers(dataPath string) ([]Datacenter, error) {
	var res []Datacenter

	dataBytes, err := ioutil.ReadFile(dataPath)
	if err != nil {
		return res, err
	}

	err = yaml.Unmarshal(dataBytes, &res)
	if err != nil {
		return res, err
	}

	return res, nil
}

func updateRack(rack *Rack, version int) bool {
	// turning off rack grafecully
	rack.Status = "offline"

	// waitGroup for goroutines
	var wg sync.WaitGroup
	wg.Add(len(rack.Servers))

	for i := 0; i < len(rack.Servers); i++ {
		go func(server *Server) {
			defer wg.Done()
			oldVersion := server.SoftwareVersion

			server.SoftwareVersion = version
			fmt.Printf("server %v was updated from %v to %v\n", server.Name, oldVersion, server.SoftwareVersion)
		}(&rack.Servers[i])
	}

	wg.Wait()
	fmt.Printf("Updated rack: %v , status: %v, servers : %v\n\n", rack.Name, rack.Status, rack.Servers)
	return true
}

func main() {
	testInputFile := "servers_input.yaml"
	newSoftwareVersion := 2

	dcs, err := getServers(testInputFile)
	if err != nil {
		panic(err)
	}
	fmt.Printf("All DC's before update: %v\n\n\n", dcs)

	for _, dc := range dcs {
		fmt.Printf("Here is dc %v, it contains %v racks for %v servers, updating it....\n", dc.Name, len(dc.Racks), len(dc.Racks[0].Servers))
		for _, r := range dc.Racks {
			updated := updateRack(&r, newSoftwareVersion)
			if !updated {
				panic("Failed to update rack")
			}

			r.Status = "online"
		}
	}

	fmt.Printf("All DC's updated: %v\n", dcs)
}
