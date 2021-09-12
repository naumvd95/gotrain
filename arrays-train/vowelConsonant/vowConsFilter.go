package main

import "fmt"

// time: O(n) | space: O(n)
var vowelsMap = map[rune]bool{
	'a': true,
	'b': false,
	'c': false,
	'd': false,
	'e': true,
	'f': false,
	'g': false,
	'h': false,
	'i': true,
	'j': false,
	'k': false,
	'l': false,
	'm': false,
	'n': false,
	'o': true,
	'p': false,
	'q': false,
	'r': false,
	's': false,
	't': false,
	'u': true,
	'v': false,
	'w': false,
	'x': false,
	'y': true,
	'z': false,
}

// return string w/o a vowels
func filterVowels(s string) string {
	res := ""
	for _, c := range s {
		if isVowel := vowelsMap[c]; !isVowel {
			res += string(c)
		}
	}
	return res
}

// return string w/o a vowels
func filterConsonants(s string) string {
	res := ""
	for _, c := range s {
		if isVowel := vowelsMap[c]; isVowel {
			res += string(c)
		}
	}
	return res
}

func main() {
	testData := []string{"consonants", "vowels"}

	for _, v := range testData {
		fmt.Printf("Orig: %v\nno-vowels: %v\nno-consonants: %v\n", v, filterVowels(v), filterConsonants(v))
	}
}
