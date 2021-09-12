package programmerstring

import (
	"fmt"
	"testing"
)

type data struct {
	InputString    string
	DesiredWord    string
	ResultIndicies int
}

var testData = []data{
	data{
		InputString:    "progxrammerrxproxgrammer",
		DesiredWord:    "programmer",
		ResultIndicies: 2,
	},
	data{
		InputString:    "xprogxrmaxemrppprmmograeiruu",
		DesiredWord:    "programmer",
		ResultIndicies: 2,
	},
}

func TestSubarraysSum(t *testing.T) {

	for _, v := range testData {

		res := lengthBetweenWords(v.DesiredWord, v.InputString)
		if res != v.ResultIndicies {
			t.Errorf("Expected %v indicies in between occurencies of the word %v in string %v, found: %v\n", v.ResultIndicies, v.DesiredWord, v.InputString, res)
		} else {
			fmt.Printf("HURRAY: Expected %v indicies in between occurencies of the word %v in string %v, found: %v\n", v.ResultIndicies, v.DesiredWord, v.InputString, res)
		}
	}
}
