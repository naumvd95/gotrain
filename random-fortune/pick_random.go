package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"
)

func getRandomByte(size int64) int64 {
	rand.Seed(time.Now().UnixNano())
	return rand.Int63n(size)
}

func pickRandomLine(offset int64, file *os.File) (string, error) {
	res := []byte{}

	// seekerSlice is single-item slice to perform one-by-one reading
	seekerSlice := make([]byte, 1)
	rightTraverser := offset
	// from the offset to the right until reaching endline
	for seekerSlice[0] != '\n' {
		_, err := file.ReadAt(seekerSlice, rightTraverser)
		if err != nil {
			return string(res), err
		}

		rightTraverser++
		res = append(res, seekerSlice[0])
	}

	seekerSlice = make([]byte, 1)
	leftTraverser := offset - 1 // we already got bytes from offset var
	// from the offset to the right until *nothing to read* OR
	// reaching start of the file OR
	// reaching endline of the previous line
	for leftTraverser >= 0 && seekerSlice[0] != '\n' {
		bytesRead, err := file.ReadAt(seekerSlice, leftTraverser)
		if err != nil {
			return string(res), err
		}
		// there is nothing to read else TODO use it in above as well ^^
		if bytesRead == 0 {
			break
		}

		leftTraverser--
		res = append(seekerSlice, res...)
	}

	return string(res), nil
}

func main() {
	// just open w/o reading all
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	// getting file stat to get file size
	fstat, err := file.Stat()
	if err != nil {
		panic(err)
	}

	offset := getRandomByte(fstat.Size())
	fmt.Printf("random offset: %v\n", offset)

	randomLine, err := pickRandomLine(offset, file)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Hey, rand fortune is: %v\n", randomLine)
}
