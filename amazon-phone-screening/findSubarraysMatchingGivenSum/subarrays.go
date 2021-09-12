package subarrayssum

// time: O(n) | space: O(n) , where n - len of nums
func subarraySum(nums []int, k int) int {
	var result int
	prevSums := make(map[int]int) // key: sum of the subarray elements, value: amount of known subarrays matching desired sum

	currentSum := 0
	for i := 0; i < len(nums); i++ {
		currentSum += nums[i]

		if currentSum == k {
			result++
		}

		if val, exists := prevSums[currentSum-k]; exists {
			result += val
		}

		prevSums[currentSum]++
	}

	return result
}
