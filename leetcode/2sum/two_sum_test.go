package main

import (
	"fmt"
	"reflect"
	"sort"
	"testing"
)

/*
 Input: nums = [2,7,11,15], target = 9
 Output: [0,1]
 Output: Because nums[0] + nums[1] == 9, we return [0, 1].
*/

type testData struct {
	InputNumbers    []int
	TargetSum       int
	OutputPositions []int
}

var data = []testData{
	testData{
		InputNumbers:    []int{2, 7, 11, 15},
		TargetSum:       9,
		OutputPositions: []int{0, 1},
	},
	testData{
		InputNumbers:    []int{3, 2, 3},
		TargetSum:       6,
		OutputPositions: []int{0, 2},
	},
	testData{
		InputNumbers:    []int{3, 2, 4},
		TargetSum:       6,
		OutputPositions: []int{1, 2},
	},
	testData{
		InputNumbers:    []int{3, 3},
		TargetSum:       6,
		OutputPositions: []int{0, 1},
	},
}

func TestFindTwoSumIndicies(t *testing.T) {
	for _, v := range data {
		res := findTwoSumIndicies(v.InputNumbers, v.TargetSum)
		sort.Ints(res)
		if !reflect.DeepEqual(res, v.OutputPositions) {
			t.Errorf("Expected: %v, got %v\n", v.OutputPositions, res)
		} else {
			fmt.Printf("Input: %v, target: %v, expected: %v, got %v !\n", v.InputNumbers, v.TargetSum, v.OutputPositions, res)
		}
	}
}
