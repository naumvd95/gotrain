package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

/*
 * Complete the 'bfs' function below.
 *
 * The function is expected to return an INTEGER_ARRAY.
 * The function accepts following parameters:
 *  1. INTEGER n
 *  2. INTEGER m
 *  3. 2D_INTEGER_ARRAY edges
 *  4. INTEGER s
 */

type Queue []int32

func (q *Queue) isEmpty() bool {
	return len(*q) == 0
}

func (q *Queue) push(value int32) {
	*q = append(*q, value)
}

func (q *Queue) pop() (int32, bool) {
	var value int32
	if q.isEmpty() {
		return value, false
	}

	value = (*q)[0]
	*q = (*q)[1:] // cutoff

	return value, true
}

func bfs(n int32, m int32, edges [][]int32, s int32) []int32 {
	// get neighborsMap (vertex: []itsNeighbors)
	// with help of queue -> combine distancesMap(vertex: its-distance-from-start-vertex)

	neighborsMap := make(map[int32][]int32)
	for _, e := range edges {
		// mirror each append, i.e. 1->2 == 2->1
		neighborsMap[e[0]] = append(neighborsMap[e[0]], e[1])
		neighborsMap[e[1]] = append(neighborsMap[e[1]], e[0])
	}

	// init queue and start vertex
	distancesMap := make(map[int32]int32)
	distancesMap[s] = 0 // start point
	var q Queue
	q.push(s) // init start in queue as well

	for !q.isEmpty() {
		// we need to go through every neighbor that is not in distancesMap

		currentVertex, _ := q.pop()
		neighbors := neighborsMap[currentVertex]

		for _, nb := range neighbors {
			if _, exist := distancesMap[nb]; exist {
				// we already know distance, skipping it
				continue
			}
			// otherwise calculate distance according to current 'depth' from start
			distancesMap[nb] = distancesMap[currentVertex] + 1
			// and add it in queue to check its neighbors in future
			q.push(nb)
		}
	}
	// we got distancesMap with ALL distances from S to any vertex

	results := []int32{}
	// restrictions: orderd number (from 1 to n)
	// no need to print distances from Start to Start
	// if node unreachable == print -1

	for i := int32(1); i <= n; i++ {
		if i == s {
			// we dont need to report start<>start distance
			continue
		}

		if value, exist := distancesMap[i]; exist {
			results = append(results, value*6) // weight is 6
		} else {
			// vertex unreachable
			results = append(results, -1)
		}
	}

	return results
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 16*1024*1024)

	qTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)
	q := int32(qTemp)

	for qItr := 0; qItr < int(q); qItr++ {
		firstMultipleInput := strings.Split(strings.TrimSpace(readLine(reader)), " ")

		nTemp, err := strconv.ParseInt(firstMultipleInput[0], 10, 64)
		checkError(err)
		n := int32(nTemp)

		mTemp, err := strconv.ParseInt(firstMultipleInput[1], 10, 64)
		checkError(err)
		m := int32(mTemp)

		var edges [][]int32
		for i := 0; i < int(m); i++ {
			edgesRowTemp := strings.Split(strings.TrimRight(readLine(reader), " \t\r\n"), " ")

			var edgesRow []int32
			for _, edgesRowItem := range edgesRowTemp {
				edgesItemTemp, err := strconv.ParseInt(edgesRowItem, 10, 64)
				checkError(err)
				edgesItem := int32(edgesItemTemp)
				edgesRow = append(edgesRow, edgesItem)
			}

			if len(edgesRow) != 2 {
				panic("Bad input")
			}

			edges = append(edges, edgesRow)
		}

		sTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
		checkError(err)
		s := int32(sTemp)

		result := bfs(n, m, edges, s)

		for i, resultItem := range result {
			fmt.Fprintf(writer, "%d", resultItem)

			if i != len(result)-1 {
				fmt.Fprintf(writer, " ")
			}
		}

		fmt.Fprintf(writer, "\n")
	}

	writer.Flush()
}

func readLine(reader *bufio.Reader) string {
	str, _, err := reader.ReadLine()
	if err == io.EOF {
		return ""
	}

	return strings.TrimRight(string(str), "\r\n")
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
