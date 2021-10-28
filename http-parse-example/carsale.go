package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"time"
)

type Car struct {
	Name  string `yaml:"name"`
	Model string `yaml:"model"`
}

type Center struct {
	Name string `yaml:"name"`
	Cars []Car  `yaml:"cars"`
}

func getCenters(host string, port int, endpoint string) ([]Center, error) {
	var centers []Center
	//1. check port is opened
	timeout := time.Second
	conn, err := net.DialTimeout("tcp", fmt.Sprintf("%v:%v", host, port), timeout)
	if err != nil {
		return centers, err
	}
	defer conn.Close()

	resp, err := http.Get(fmt.Sprintf("http://%v:%v/%v", host, port, endpoint))
	if err != nil {
		return centers, err
	}
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return centers, err
	}

	err = json.Unmarshal(data, &centers)
	if err != nil {
		return centers, err
	}

	return centers, nil
}
func main() {
	serverIP := "0.0.0.0"
	serverPorts := []int{8000, 8001}
	inputFile := "car_sale.json"

	for _, p := range serverPorts {
		fmt.Printf("Checking port: %v\n", p)
		data, err := getCenters(serverIP, p, inputFile)
		if err != nil {
			fmt.Printf("Something went wrong accessing http://%v:%v/%v : %v\n", serverIP, p, inputFile, err)
		} else {
			fmt.Printf("Data from http://%v:%v/%v : %v\n", serverIP, p, inputFile, data)
		}
	}
}
