package main

import (
	"fmt"
	"reflect"
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
		OutputPositions: []int{1, 2},
	},
	testData{
		InputNumbers:    []int{2, 3, 3},
		TargetSum:       6,
		OutputPositions: []int{2, 3},
	},
	testData{
		InputNumbers:    []int{2, 3, 4},
		TargetSum:       6,
		OutputPositions: []int{1, 3},
	},
	testData{
		InputNumbers:    []int{3, 3},
		TargetSum:       6,
		OutputPositions: []int{1, 2},
	},
}

func TestFindTwoSumIndiciesNg(t *testing.T) {
	for _, v := range data {
		res := findTwoSumIndiciesNg(v.InputNumbers, v.TargetSum)
		if !reflect.DeepEqual(res, v.OutputPositions) {
			t.Errorf("Expected: %v, got %v\n", v.OutputPositions, res)
		} else {
			fmt.Printf("Input: %v, target: %v, expected: %v, got %v !\n", v.InputNumbers, v.TargetSum, v.OutputPositions, res)
		}
	}
}
