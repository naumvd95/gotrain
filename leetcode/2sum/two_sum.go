package main

import "fmt"

/*
Task:

Given an array of integers nums and an integer target, return indices of the two numbers such that they add up to target.

You may assume that each input would have exactly one solution, and you may not use the same element twice.

You can return the answer in any order.

Input: nums = [2,7,11,15], target = 9
Output: [0,1]
Output: Because nums[0] + nums[1] == 9, we return [0, 1].
*/

/*
COMPLEXITY:

time: o(n) , because of loop for all nums
space: o(n) , because we store all numbers in map
*/
func findTwoSumIndicies(nums []int, target int) []int {
	verifiedNums := make(map[int]int) // key: number , value: index of that number

	for i := 0; i < len(nums); i++ {
		proposal := target - nums[i] // we just forecast if proposal exist and its suitable for condition

		// checking if proposal exist in our verified map
		if _, ok := verifiedNums[proposal]; ok { // use ok as bool verification mechanism, because we already at index==0, and all future ones will be true.
			return []int{i, verifiedNums[proposal]} // return current index and proposal index if we already know that such proposal exist
		}
		verifiedNums[nums[i]] = i // otherwise we just mark current number as visited, btw here is can be our first pair number=X and index=0(!)
	}

	return nil
}

func main() {
	nums := []int{2, 7, 11, 15}
	target := 9

	res := findTwoSumIndicies(nums, target)
	fmt.Printf("Here is original %v, where target is: %v, answer is: %v", nums, target, res)
}
