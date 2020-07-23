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

func CanILeavePlace(seats []int, seatIndex, neighbors int) bool {
	// first/last seat cant influence on neighbors
	if seatIndex == 0 || seatIndex == len(seats)-1 {
		return true
	}
	l := seats[:seatIndex]
	r := seats[seatIndex+1:]

	// if closest neighbors after shift are not siblings, its ok
	if l[len(l)-1] != r[0] {
		return true
	}
	// at least 2 siblings detected
	siblingsCounter := 2
	for i := len(l) - 1; i > 0; i-- {
		if l[i] == l[i-1] {
			siblingsCounter++
		} else {
			break
		}
	}
	for i := 0; i < len(r); i++ {
		if r[i] == r[i+1] {
			siblingsCounter++
		} else {
			break
		}
	}

	return siblingsCounter < neighbors
}

func main() {
	fmt.Printf("Here is result of %v on position %v w/ neighbors %v : %v\n", cinemaSeats, yourSeat, neighborsSiblilngs, CanILeavePlace(cinemaSeats, yourSeat, neighborsSiblilngs))
}
