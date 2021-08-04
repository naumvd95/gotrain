package main

import (
	"fmt"
	"testing"
)

type testData struct {
	NumberOfSeats  int64
	SocialDistance int64
	Dinners        int32
	DinnersSeats   []int64
	Result         int64
}

var data = []testData{
	testData{
		NumberOfSeats:  10,
		SocialDistance: 1,
		Dinners:        2,
		DinnersSeats:   []int64{2, 6},
		Result:         3,
	},
	testData{
		NumberOfSeats:  15,
		SocialDistance: 2,
		Dinners:        3,
		DinnersSeats:   []int64{11, 6, 14},
		Result:         1,
	},
	testData{
		NumberOfSeats:  100,
		SocialDistance: 2,
		Dinners:        10,
		DinnersSeats:   []int64{1, 4, 7, 10, 13, 16, 19, 22, 25, 28},
		Result:         24,
	},
}

func TestGetMaxAdditionalDinersCount(t *testing.T) {
	for _, dataSet := range data {
		fmt.Printf("Number of seats: %v, distance: %v, busy seats: %v, expected extra seats: %v\n", dataSet.NumberOfSeats, dataSet.SocialDistance, dataSet.DinnersSeats, dataSet.Result)

		actualRes := GetMaxAdditionalDinersCount(dataSet.NumberOfSeats, dataSet.SocialDistance, dataSet.Dinners, dataSet.DinnersSeats)
		if actualRes != dataSet.Result {
			t.Errorf("Expected result: %v, got: %v\n", dataSet.Result, actualRes)
		} else {
			fmt.Printf("Result: %v\n", actualRes)
		}
	}
}
