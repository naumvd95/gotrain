package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	file, err := os.Open("nginx-k8s-pod.log")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	result := make(map[int]string) // key: num of line, value: data
	lines := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines++
		line := scanner.Text()

		if !strings.Contains(line, "GET") {
			continue
		}

		if !strings.Contains(line, "credentials") {
			continue
		}
		result[lines] = line
	}

	fmt.Printf("Found %v credentials: %v\n", len(result), result)
}
