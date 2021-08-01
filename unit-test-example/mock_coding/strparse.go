package main

import (
	"encoding/csv"
	"fmt"
	"strings"
)

const (
	csvData = `
Username; Identifier;First name;Last name
booker12;9012;Rachel;Booker
grey07;2070;Laura;Grey
johnson81;4081;Craig;Johnson
jenkins46;9346;Mary;Jenkins
smith79;5079;Jamie;Smith
`
)

func main() {

	csv := csv.NewReader(strings.NewReader(csvData))
	csv.Comma = ';'
	dataLines, err := csv.ReadAll()
	if err != nil {
		panic(err)
	}

	keys := []string{}
	for _, csvKey := range dataLines[0] {
		keys = append(keys, csvKey)
	}

	users := []map[string]string{}
	for _, line := range dataLines[1:] {
		user := make(map[string]string)

		for idx, v := range line {
			user[keys[idx]] = v
		}
		users = append(users, user)
	}

	fmt.Printf("Hey, here is all users: %v\n", users)
}
