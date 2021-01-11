package main

import (
	"fmt"
	"testing"
)

type data struct {
	Input          string
	AcceptedOutput []string
}

var dataSet = []data{
	data{
		Input: "babad",
		AcceptedOutput: []string{
			"bab",
			"aba",
		},
	},
	data{
		Input: "cbbd",
		AcceptedOutput: []string{
			"bb",
		},
	},
	data{
		Input: "a",
		AcceptedOutput: []string{
			"a",
		},
	},
	data{
		Input: "ac",
		AcceptedOutput: []string{
			"a",
			"c",
		},
	},
}

func TestLongestPalindrome(t *testing.T) {
	for _, v := range dataSet {
		fmt.Printf("Longest palindromes for %v are: %v\n, checking\n", v.Input, v.AcceptedOutput)
		res := longestPalindrome(v.Input)

		isMatched := false
		for _, answer := range v.AcceptedOutput {
			if res == answer {
				isMatched = true
			}
		}

		if !isMatched {
			t.Errorf("String %v, accepted answers: %v, got %v\n", v.Input, v.AcceptedOutput, res)
		}
	}
}
