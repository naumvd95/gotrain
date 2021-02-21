package main

import "fmt"

/*
Given an array of integers that is already sorted in ascending order, find two numbers such that they add up to a specific target number.

The function twoSum should return indices of the two numbers such that they add up to the target, where index1 must be less than index2.

Note:

Your returned answers (both index1 and index2) are not zero-based.
You may assume that each input would have exactly one solution and you may not use the same element twice.

Input: numbers = [2,7,11,15], target = 9
Output: [1,2]
Explanation: The sum of 2 and 7 is 9. Therefore index1 = 1, index2 = 2.
*/

/*
COMPLEXITY:
time: O(n*logn) - "sort complexity" * "for loop until head<tail"
space: O(1) - because we dont store anything additionally

SOLUTION MINDSET:

As your array already sorted, you dont need to use 'proposal' to forecast number.
Use head/tail scenario, where you iterating via nums and combine sums of first/last element.
If sum <> deesired target, you may drive head++ or tail-- to be closer to desired sum

*/
func findTwoSumIndiciesNg(nums []int, target int) []int {
	// start/end of our slice (handle 0 index later)
	head := 1
	tail := len(nums)

	for head < tail { // handle loop w/ continue and return
		sum := nums[head-1] + nums[tail-1] // start to calculate possible sum and compare w/ target

		if sum > target { // if its too much, decrease our tail [slice sorted ASC]
			tail--
			continue
		}
		if sum < target { // if need more, increase our head [slice sorted ASC]
			head++
			continue
		}
		return []int{head, tail} // if its equals, gotcha
	}

	return nil
}

func main() {
	nums := []int{2, 7, 11, 15}
	target := 9

	res := findTwoSumIndiciesNg(nums, target)
	fmt.Printf("Advanced 2sum II: Here is original %v, where target is: %v, answer is: %v", nums, target, res)
}
