package removeandsplitarray

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
		Input:         []int{6, 2, 3, 2, 1},
		CanBeSplitted: true,
		Details: `On removing element 2 at index 1,
the array gets divided into two subarrays [6]
and [3, 2, 1] having equal sum`,
	},
	testData{
		Input:         []int{6, 1, 3, 2, 5},
		CanBeSplitted: true,
		Details: `On removing element 3 at index 2,
the array gets divided into two subarrays [6, 1]
and [2, 5] having equal sum`,
	},
	testData{
		Input:         []int{6, -2, -3, 2, 3},
		CanBeSplitted: true,
		Details: `On removing element 6 at index 0, 
the array gets divided into two sets []
and [-2, -3, 2, 3] having equal sum`,
	},
	testData{
		Input:         []int{6, -2, 3, 2, 3},
		CanBeSplitted: false,
		Details:       "cannot-be-splitted",
	},
}

func TestSplitBySum(t *testing.T) {
	for _, v := range input {
		result := removeAndSplit(v.Input)
		if result != v.CanBeSplitted {
			t.Fatalf("%v (answer: %v),  but result was: %v\n\n", v.Input, v.Details, result)
		} else {
			fmt.Printf("%v (answer: %v) test result was: %v\n\n", v.Input, v.Details, result)
		}
	}
}
