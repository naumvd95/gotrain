package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

/*Algo

1. parse file line-by-line (file size may be too big to load in-memory) | time: O(n) | Space: O(m), where n - number of all lines, m - max length of line
2. Make a map from all lines, (key: name, value: itsNumber) -> getting rid of the duplicates | time: O(n) | space: O(u), where u - number of unique keys
3. Convert map into slice of slices (each sub slice is [a,b], where a: name, b: itsNumber) | time: O(u) | space: O(u), where u - number of unique keys
4. Sort slice by the 2nd value | quickSort time: u*log(u) | space: O(u), where u - number of unique keys

Overall worst complexity:
time: O(u*log(u)), where u - number of unique keys
space: O(u), where u - number of unique keys
*/

type Person struct {
	Name   string
	Number int
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	//1. scan file line by line and 2. generate map
	uniqNamesMap := make(map[string]int)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		kvStr := strings.Split(line, "=")
		name := kvStr[0]
		num, err := strconv.Atoi(kvStr[1])
		if err != nil {
			panic(err)
		}

		uniqNamesMap[name] = num
	}

	//3. convert map into slice
	uniqNamesSlice := []Person{}
	for k, v := range uniqNamesMap {
		uniqNamesSlice = append(uniqNamesSlice, Person{
			Name:   k,
			Number: v,
		})
	}

	//4. sort slice
	sort.Slice(uniqNamesSlice, func(a, b int) bool {
		return uniqNamesSlice[a].Number < uniqNamesSlice[b].Number
	})
	fmt.Printf("Here is sorted dataL %v\n", uniqNamesSlice)
}
