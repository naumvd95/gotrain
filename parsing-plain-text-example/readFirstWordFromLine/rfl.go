package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func readFirstword(s string) string {
	arr := strings.Split(s, " ")

	return arr[0]
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	sc := bufio.NewScanner(file)
	for sc.Scan() {
		line := sc.Text()
		fmt.Printf("Here is: %v\n", readFirstword(line))
	}
}
