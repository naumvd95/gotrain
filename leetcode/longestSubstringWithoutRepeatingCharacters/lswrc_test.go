package main

import (
	"fmt"
	"testing"
)

type data struct {
	Input  string
	Output int
}

var dataSet = []data{
	data{
		Input:  "abcabcbb",
		Output: 3,
	},
	data{
		Input:  "bbbbb",
		Output: 1,
	},
	data{
		Input:  "pwwkew",
		Output: 3,
	},
	data{
		Input:  "",
		Output: 0,
	},
}

func TestLengthOfLongestSubstring(t *testing.T) {
	for _, v := range dataSet {
		res := lengthOfLongestSubstring(v.Input)
		fmt.Printf("testcase %v(%v), result: %v\n", v.Input, v.Output, res)
		if res != v.Output {
			t.Errorf("Incorrect answer %v, expected %v\n", res, v.Output)
		}
	}
}
