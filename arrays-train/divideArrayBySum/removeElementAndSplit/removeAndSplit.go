package removeandsplitarray

import "fmt"

/* algo:
1. loop over array and find total sum
2. loop over arral and find right&left sums for each element:
  - leftSum can be managed as sum_so_far during traversing
  - rightSum can be managed as (totalSum - sum_so_far - current element) during traversing
  - dont forget to adjust sum_so_far only after all calculations
3. if leftSum == rightSum >> true

time: O(n) | space: O(1)
*/
func removeAndSplit(array []int) bool {
	//1.
	var totalSum int
	for _, v := range array {
		totalSum += v
	}

	var (
		leftSum  int
		rightSum int
		sumSoFar int
	)

	for idx, v := range array {
		leftSum = sumSoFar
		rightSum = totalSum - sumSoFar - v

		if leftSum == rightSum {
			fmt.Printf("Hurray, array may be splitted in %v and %v by removing element %v at the idx %v\n\n", array[:idx], array[idx+1:], v, idx)
			return true
		}

		sumSoFar += v
	}

	return false
}
