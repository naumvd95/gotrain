package main

import (
	"fmt"
	"testing"
)

type testData struct {
	AllSeats  []int
	Neighbors int
	YourSeat  int
	Result    bool
}

var data = []testData{
	testData{
		AllSeats:  []int{3, 45, 4, 4, 4, 5, 4, 1, 2, 5, 6},
		Neighbors: 3,
		YourSeat:  5,
		Result:    false,
	},
	testData{
		AllSeats:  []int{3, 2, 5, 6},
		Neighbors: 3,
		YourSeat:  3,
		Result:    true,
	},
	testData{
		AllSeats:  []int{3, 2, 5, 6},
		Neighbors: 2,
		YourSeat:  0,
		Result:    true,
	},
	testData{
		AllSeats:  []int{3, 4, 5, 5, 5, 5, 6},
		Neighbors: 2,
		YourSeat:  3,
		Result:    false,
	},
	testData{
		AllSeats:  []int{3, 4, 5, 5, 5, 5, 6},
		Neighbors: 2,
		YourSeat:  6,
		Result:    true,
	},
}

func TestCanILeavePlace(t *testing.T) {
	msg := "Array: %v\n Neighbors: %v\n YourSeat: %v\n answer: %v\n-------------------------\n"

	for _, dataSet := range data {
		fmt.Printf(msg, dataSet.AllSeats, dataSet.Neighbors, dataSet.YourSeat, dataSet.Result)
		actualRes := CanILeavePlace(dataSet.AllSeats, dataSet.YourSeat, dataSet.Neighbors)
		if actualRes != dataSet.Result {
			t.Errorf("Expected result: %v, got: %v\n", dataSet.Result, actualRes)
		}
	}
}
