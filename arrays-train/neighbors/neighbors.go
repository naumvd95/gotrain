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

// time: O(n) | space: O(1)
func CanILeavePlace(seats []int, seatIndex, maxSiblings int) bool {
	if seatIndex == 0 || seatIndex == len(seats)-1 {
		// there is no changes if you last or first
		return true
	}

	// WARNING: space: O(n) iterate via original array to reduce space complexity
	leftPart := seats[:seatIndex]    // exсluding seatIndex
	rightPart := seats[seatIndex+1:] // excluding seatIndex

	if leftPart[len(leftPart)-1] != rightPart[0] {
		// if closest neighbors are different, there is nothing to calculate else
		return true
	}
	siblingsCounter := 2             // because of the above condition
	siblingToCompare := rightPart[0] // choose one of 2 siblings for futher comparing

	for i := len(leftPart) - 2; i >= 0; i-- {
		if leftPart[i] == siblingToCompare {
			siblingsCounter++
			if siblingsCounter >= maxSiblings {
				return false
			}
		}
	}
	for i := 1; i < len(rightPart); i++ {
		if rightPart[i] == siblingToCompare {
			siblingsCounter++
			if siblingsCounter >= maxSiblings {
				return false
			}
		}
	}

	return true
}

func main() {
	fmt.Printf("Here is result of %v on position %v w/ neighbors %v : %v\n", cinemaSeats, yourSeat, neighborsSiblilngs, CanILeavePlace(cinemaSeats, yourSeat, neighborsSiblilngs))
}
