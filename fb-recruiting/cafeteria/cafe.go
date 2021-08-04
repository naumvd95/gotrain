package main

import "sort"

// time: O(d logd) where d = dinners | space: O(1)
func GetMaxAdditionalDinersCount(numberOfSeats int64, socialDistance int64, dinners int32, dinnersSeats []int64) int64 {
	/* Algo:
	    We cant resolve this task by building whole table of numberOfSeats, because numberOfSeats is up to 10^15 , so we can go with linear solution
		1. Sort dinnersSeats ASC , to understand indexes of seats: O(mlogm) where m is dinners
		2. Iterate via each segment in-between 2 dinnersSeats, and use following formula:
		   - extraSeats := (rightS - leftS - socialDistance - 1) // (socialDistance + 1), where:
		     rightS = i in dinnersSeats
			 leftS = i+1 in dinnersSeats
			 (rightS - leftS - socialDistance - 1) = socialDistance coz we need to decrement at least 1 distance before iterating, and -1 needed to get pure amount of seats from rightS - leftS
			 // (socialDistance + 1) == divide w/o leftover == its a place needed for 1 more person
		3. handle edge case like: from 1 to leftFirst && from rightLast to N, because 1st and last places may be also a free seats + in-between such segments there may be free seats:
		   - extraSeatsStart := (dinnersSeats[0] - 1) // (socialDistance + 1)
		   - extraSeatsEnd := (numberOfSeats - dinnersSeats[dinners-1]) // (socialDistance + 1)
		4. sum up all extra seats
	*/
	var extraSeats int64

	// 1. time: O(d logd) where d = dinners | space: O(1)
	sort.Slice(dinnersSeats, func(i, j int) bool {
		return dinnersSeats[i] < dinnersSeats[j]
	})

	//2. time: O(dinners) | space: O(1)
	for i := 0; i < len(dinnersSeats)-1; i++ {
		leftS := dinnersSeats[i]
		rightS := dinnersSeats[i+1]

		extraSeats += (rightS - leftS - socialDistance - 1) / (socialDistance + 1)
	}

	//3. time: O(1) | space: O(1)
	extraSeatsStart := (dinnersSeats[0] - 1) / (socialDistance + 1)
	extraSeatsEnd := (numberOfSeats - dinnersSeats[dinners-1]) / (socialDistance + 1)

	//4. time: O(1) | space: O(1)
	extraSeats += extraSeatsStart + extraSeatsEnd

	return extraSeats
}
