package splitarray

import (
	"fmt"
	"testing"
)

type testData struct {
	Input         []int
	CanBeSplitted bool
	Details       string
}

var input = []testData{
	testData{
		Input:         []int{1, 3, 1, 1, 1, 2, 1, 2},
		CanBeSplitted: true,
		Details:       "[1, 3, 1, 1,| 1, 2, 1, 2]",
	},
	testData{
		Input:         []int{1, 3, 1, 1, 1, 2, 1, 3},
		CanBeSplitted: false,
		Details:       "[cannot-be-splitted]",
	},
	testData{
		Input:         []int{1, 3, 1, 1, 1, 2, 1, 4},
		CanBeSplitted: true,
		Details:       "[1, 3, 1, 1, 1,| 2, 1, 4]",
	},
}

func TestSplitBySum(t *testing.T) {
	for _, v := range input {
		result := splitBySum(v.Input)
		if result != v.CanBeSplitted {
			t.Fatalf("%v (answer: %v),  but result was: %v\n", v.Input, v.Details, result)
		} else {
			fmt.Printf("%v (answer: %v) test result was: %v\n", v.Input, v.Details, result)
		}
	}
}
