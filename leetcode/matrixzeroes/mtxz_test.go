package main

import (
	"reflect"
	"testing"
)

type data struct {
	Input         [][]int
	DesiredResult [][]int
}

var dataSet = []data{
	data{
		Input: [][]int{
			[]int{1, 1, 1},
			[]int{1, 0, 1},
			[]int{1, 1, 1},
		},
		DesiredResult: [][]int{
			[]int{1, 0, 1},
			[]int{0, 0, 0},
			[]int{1, 0, 1},
		},
	},
	data{
		Input: [][]int{
			[]int{0, 1, 2, 0},
			[]int{3, 4, 5, 2},
			[]int{1, 3, 1, 5},
		},
		DesiredResult: [][]int{
			[]int{0, 0, 0, 0},
			[]int{0, 4, 5, 0},
			[]int{0, 3, 1, 0},
		},
	},
	data{
		Input: [][]int{
			[]int{-4, -2147483648, 6, -7, 0},
			[]int{-8, 6, -8, -6, 0},
			[]int{2147483647, 2, -9, -6, 10},
		},
		DesiredResult: [][]int{
			[]int{0, 0, 0, 0, 0},
			[]int{0, 0, 0, 0, 0},
			[]int{2147483647, 2, -9, -6, 0},
		},
	},
}

func TestSetZeroes(t *testing.T) {
	for _, v := range dataSet {
		res := v.Input
		setZeroes(res)
		if !reflect.DeepEqual(res, v.DesiredResult) {
			t.Errorf("Matrix: %v\n Expected: %v\n, got %v\n", v.Input, v.DesiredResult, res)
		}
	}
}
