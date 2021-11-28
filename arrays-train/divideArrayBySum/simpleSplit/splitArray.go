package splitarray

import "fmt"

func splitBySum(a []int) bool {
	/*
		algo:
		1. calculate whole sum of array elements and divide by 2 == halfSum
		2. loop over array, and if you reach desired halfSum == means that here we can divide
		3. if totalSum cannot be divided by 2 == false!
	*/

	var totalSum int
	for _, v := range a {
		totalSum += v
	}
	if totalSum%2 != 0 {
		return false
	}

	halfSum := totalSum / 2
	var leftArraySum int
	for _, v := range a {
		leftArraySum += v
		if leftArraySum == halfSum {
			return true
		}
	}

	return false
}

func main() {
	input := []int{1, 3, 1, 1, 1, 2, 1, 2}

	canBeSplitted := splitBySum(input)
	if canBeSplitted {
		fmt.Printf("array %v can be splitted\n", input)
	} else {
		fmt.Printf("array %v can NOT be splitted\n", input)
	}
}
