package main

/*
Imagine:
- You are at cinema, sitting on seat aSeat in array allSeats == allSeats[aSeat]
- If you went out, your neighbors joined to each other
- You need to check that after join N members are not equal , i.e. members are not siblings for example
*/
import "fmt"

var (
	cinemaSeats        = []int{3, 45, 4, 4, 4, 5, 4, 1, 2, 5, 6}
	neighborsSiblilngs = 3
	yourSeat           = 5
)

func CanILeavePlace(seats []int, seatIndex, maxSiblings int) bool {
	if seatIndex == 0 || seatIndex == len(seats)-1 {
		// if u last or first, there is no matter
		return true
	}

	l := seatIndex - 1
	r := seatIndex + 1
	if seats[l] != seats[r] {
		// nothing was changed, because your closest are not siblings
		return true
	}

	// otherwise we already have 2 siblings
	// choose one to compare
	goldenNeighbor := seats[l]
	siblingCounter := 2
	l--
	r++

	// traverse to left
	for l >= 0 {
		if seats[l] != goldenNeighbor {
			break
		}

		siblingCounter++
		l--

		if siblingCounter > maxSiblings {
			return false
		}
	}

	// traverse to right
	for r < len(seats) {
		if seats[r] != goldenNeighbor {
			break
		}

		siblingCounter++
		r++

		if siblingCounter > maxSiblings {
			return false
		}
	}

	return siblingCounter < maxSiblings
}

func main() {
	fmt.Printf("Here is result of %v on position %v w/ neighbors %v : %v\n", cinemaSeats, yourSeat, neighborsSiblilngs, CanILeavePlace(cinemaSeats, yourSeat, neighborsSiblilngs))
}
