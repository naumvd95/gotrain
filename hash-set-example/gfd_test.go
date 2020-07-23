package main

import (
	"testing"
)

var SetResMap = make(map[int][]int)

func init() {
	SetResMap[7] = []int{1, 4, 6, 7, 7, 2, 3, 5, 9, 434, 834, 2}
	SetResMap[3] = []int{2, 1, 3, 5, 3, 2}
	SetResMap[-1] = []int{1}
	SetResMap[1] = []int{1, 1, 2, 2, 1}
}

func TestGetFirstDuplicate(t *testing.T) {
	var testedRes int
	msg := "Incorrect test result on dataset %v, wanted %v, get %v"
	for desiredRes, dataSet := range SetResMap {
		testedRes = GetFirstDuplicate(dataSet)
		if testedRes != desiredRes {
			t.Errorf(msg, dataSet, desiredRes, testedRes)
		}
	}
}
