package subarrayssum

import (
	"fmt"
	"testing"
)

type data struct {
	InputArray      []int
	DesiredSum      int
	ResultSubarrays int
}

var testData = []data{
	data{
		InputArray:      []int{10, 2, -2, -20, 10},
		DesiredSum:      -10,
		ResultSubarrays: 3,
	},
	data{
		InputArray:      []int{9, 4, 20, 3, 10, 5},
		DesiredSum:      33,
		ResultSubarrays: 2,
	},
	data{
		InputArray:      []int{1, 1, 1},
		DesiredSum:      2,
		ResultSubarrays: 2,
	},
	data{
		InputArray:      []int{1, 2, 3},
		DesiredSum:      3,
		ResultSubarrays: 2,
	},
}

func TestSubarraysSum(t *testing.T) {

	for _, v := range testData {

		res := subarraySum(v.InputArray, v.DesiredSum)
		if res != v.ResultSubarrays {
			t.Errorf("Expected %v subarrays matching sum %v in array %v, found: %v\n", v.ResultSubarrays, v.DesiredSum, v.InputArray, res)
		} else {
			fmt.Printf("HURRAY: Expected %v subarrays matching sum %v in array %v, found: %v\n", v.ResultSubarrays, v.DesiredSum, v.InputArray, res)
		}
	}
}
