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
 * Complete the 'quickestWayUp' function below.
 *
 * The function is expected to return an INTEGER.
 * The function accepts following parameters:
 *  1. 2D_INTEGER_ARRAY ladders
 *  2. 2D_INTEGER_ARRAY snakes
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

func quickestWayUp(ladders [][]int32, snakes [][]int32) int32 {
	/* BFS way
	   1. Generate snake/ladders map key: start, value: end
	   2. generate neighborsMap (key: cell, value: [] of all possible ways from that cell)
	   neighborsMap should follow next conditions:
	   - overall u should have 100 keys (number of cells)
	   - value for each cell will be equal from n+1 to n+6 , because of cube edges
	   - BUT u need to keep in mind that if during path search u reaching snake/ladder
	   u should add path from the end of snake/ladder like it just teleports you
	   3. Use Queue and BFS to generate distancesMap
	     (key: cell, value: number-of-cube-throws from start cell)
	   4. distancesMap[100] == is an answer, otherwise = print -1, i.e. there is no way
	      to reach finish
	*/

	//1. snake/ladders maps
	snakeMap := make(map[int32]int32)
	for _, snakePath := range snakes {
		start := snakePath[0]
		end := snakePath[1]
		snakeMap[start] = end
	}
	ladderMap := make(map[int32]int32)
	for _, ladderPath := range ladders {
		start := ladderPath[0]
		end := ladderPath[1]
		ladderMap[start] = end
	}

	// 2. get neighborsMap
	neighborsMap := make(map[int32][]int32)
	for cell := int32(1); cell <= 100; cell++ {
		nextCell := cell + 1
		cubeThrowCounter := 1
		for nextCell <= 100 && cubeThrowCounter <= 6 {
			neighbor := int32(nextCell)

			// check teleport cases for snake
			if snakeTail, exist := snakeMap[neighbor]; exist {
				neighbor = snakeTail
			}
			// check teleport cases for ladder
			if ladderTop, exist := ladderMap[neighbor]; exist {
				neighbor = ladderTop
			}

			neighborsMap[cell] = append(neighborsMap[cell], neighbor)
			nextCell++
			cubeThrowCounter++
		}
	}

	// 3. init distancesMap and queue
	distancesMap := make(map[int32]int32)
	distancesMap[int32(1)] = int32(0)
	var q Queue
	q.push(1)
	for !q.isEmpty() {
		currentCell, _ := q.pop() // no need to empty-check , covered w/ 'for' condition

		neighbors := neighborsMap[currentCell]
		for _, nb := range neighbors {
			if _, wayExist := distancesMap[nb]; wayExist {
				// skipping it, we already know path
				continue
			}
			// otherwise calculating and add to a queue to check its neighbors later
			distancesMap[nb] = distancesMap[currentCell] + 1 // add 1 more cube throw
			q.push(nb)
		}
	}

	if wayToFinish, exist := distancesMap[int32(100)]; exist {
		return wayToFinish
	}
	return -1
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 16*1024*1024)

	tTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)
	t := int32(tTemp)

	for tItr := 0; tItr < int(t); tItr++ {
		nTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
		checkError(err)
		n := int32(nTemp)

		var ladders [][]int32
		for i := 0; i < int(n); i++ {
			laddersRowTemp := strings.Split(strings.TrimRight(readLine(reader), " \t\r\n"), " ")

			var laddersRow []int32
			for _, laddersRowItem := range laddersRowTemp {
				laddersItemTemp, err := strconv.ParseInt(laddersRowItem, 10, 64)
				checkError(err)
				laddersItem := int32(laddersItemTemp)
				laddersRow = append(laddersRow, laddersItem)
			}

			if len(laddersRow) != 2 {
				panic("Bad input")
			}

			ladders = append(ladders, laddersRow)
		}

		mTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
		checkError(err)
		m := int32(mTemp)

		var snakes [][]int32
		for i := 0; i < int(m); i++ {
			snakesRowTemp := strings.Split(strings.TrimRight(readLine(reader), " \t\r\n"), " ")

			var snakesRow []int32
			for _, snakesRowItem := range snakesRowTemp {
				snakesItemTemp, err := strconv.ParseInt(snakesRowItem, 10, 64)
				checkError(err)
				snakesItem := int32(snakesItemTemp)
				snakesRow = append(snakesRow, snakesItem)
			}

			if len(snakesRow) != 2 {
				panic("Bad input")
			}

			snakes = append(snakes, snakesRow)
		}

		result := quickestWayUp(ladders, snakes)

		fmt.Fprintf(writer, "%d\n", result)
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
