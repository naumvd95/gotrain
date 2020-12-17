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

 Input
[-2,0,1,1,2]
Output
[[-2,0,2]]
Expected
[[-2,0,2],[-2,1,1]]
*/

type testData struct {
	InputNumbers    []int
	OutputPositions [][]int
}

var data = []testData{
	testData{
		InputNumbers: []int{-1, 0, 1, 2, -1, -4},
		OutputPositions: [][]int{
			[]int{-1, -1, 2},
			[]int{-1, 0, 1},
		},
	},
	testData{
		InputNumbers: []int{-2, 0, 1, 1, 2},
		OutputPositions: [][]int{
			[]int{-2, 0, 2},
			[]int{-2, 1, 1},
		},
	},
}

func TestThreeSumPrimitive(t *testing.T) {
	for _, v := range data {
		res := threeSumPrimitive(v.InputNumbers)

		for _, desiredResSlice := range v.OutputPositions {
			match := false

			for _, actualResSlice := range res {
				if reflect.DeepEqual(desiredResSlice, actualResSlice) {
					match = true
				}
			}
			if !match {
				t.Errorf("Expected: %v, got %v\n", v.OutputPositions, res)
			}
		}
		fmt.Printf("Input: %v, expected: %v, got %v !\n", v.InputNumbers, v.OutputPositions, res)
	}
}

func TestThreeSumFast(t *testing.T) {
	for _, v := range data {
		res := threeSumFast(v.InputNumbers)

		for _, desiredResSlice := range v.OutputPositions {
			match := false

			for _, actualResSlice := range res {
				if reflect.DeepEqual(desiredResSlice, actualResSlice) {
					match = true
				}
			}
			if !match {
				t.Errorf("Expected: %v, got %v\n", v.OutputPositions, res)
			}
		}
		fmt.Printf("Input: %v, expected: %v, got %v !\n", v.InputNumbers, v.OutputPositions, res)
	}
}
