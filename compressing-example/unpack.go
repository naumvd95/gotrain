package main

import (
	"errors"
	"fmt"
	"strconv"
	"unicode"
)

// time: O(n*m) , where m == sum of all multiplicators | space: O(n*m)
// if quota limit 50 ->> O(1) for space & time
func decodeStr(s string) (string, error) {
	res := ""
	counter := "" // base for our multiplicator

	for i := 0; i < len(s); i++ {
		symbol := rune(s[i])

		if unicode.IsLetter(symbol) {
			slice := []rune{symbol}

			if counter != "" {
				multiplicator, err := strconv.Atoi(counter)
				if err != nil {
					return res, err
				}

				// quota check
				if len(res)+multiplicator > 50 {
					return res, errors.New("Fail: 50 symbols quota reached")
				}

				//  c = 1 coz we already have 1 symbol
				for c := 1; c < multiplicator; c++ {
					slice = append(slice, symbol)
				}
			}
			res += string(slice)
			// reset counter
			counter = ""
		} else {
			counter += string(symbol)
		}

		// quota check
		if len(res) > 50 {
			return res, errors.New("Fail: 50 symbols quota reached")
		}
	}

	// if last symbol is a number
	if counter != "" {
		res += counter
	}

	return res, nil
}

func main() {
	testStr := "4A10BD3C"
	testFailStr := "10A10B10C20D6E"

	encOne, err := decodeStr(testStr)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Str %v encoded %v\n", testStr, encOne)

	encTwo, err := decodeStr(testFailStr)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Str %v encoded %v\n", testFailStr, encTwo)
}
