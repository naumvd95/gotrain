package main

import (
	"fmt"
	"sort"
)

/*
Given an array nums of n integers, are there elements a, b, c in nums such that a + b + c = 0? Find all unique triplets in the array which gives the sum of zero.
Notice that the solution set must not contain duplicate triplets.

Example 1:

Input: nums = [-1,0,1,2,-1,-4]
Output: [[-1,-1,2],[-1,0,1]]
Example 2:

Input: nums = []
Output: []
Example 3:

Input: nums = [0]
Output: []

Constraints:
0 <= nums.length <= 3000
-105 <= nums[i] <= 105
*/

type hashSet struct{} // weight less than bool

/* threeSumPrimitive solution w/o sorting whole array. Its more primitive way,
   because anyway we need to sort it under the hood to create unique keys for hashMap
   TODO WHATS about complexity here???
*/
func threeSumPrimitive(nums []int) [][]int {
	res := [][]int{} // final result

	if len(nums) < 3 {
		return res
	}

	if len(nums) == 3 {
		if nums[0]+nums[1]+nums[2] == 0 {
			return append(res, nums)
		}
		return res
	}

	var isMember hashSet
	hashMapTriplets := make(map[[3]int]hashSet) // hash map for keeping only unique triplets!
	// A + B + C = 0 is our current expression
	for i := 0; i < len(nums)-1; i++ { // limit range because of nested loop for sum2

		twoNumExpressionTarget := nums[i]          // declare A, now we have only B and C as typical 2sum problem
		twoSumVerifiedMap := make(map[int]hashSet) // store all numbers that passed proposal verification

		for j := i + 1; j < len(nums); j++ { // sum2
			// here we need to find C, thats meet expression  -(A + B) = C , because only in such case we expect:  A + B + C = 0
			proposalNum := -1 * (twoNumExpressionTarget + nums[j])

			if _, ok := twoSumVerifiedMap[proposalNum]; ok { // check if proposal is real, and we met them earlier
				// if so, we need to prepare hashArray to compare , to prevent result duplicates
				tempSlice := []int{proposalNum, twoNumExpressionTarget, nums[j]}
				sort.Ints(tempSlice)
				hashSlice := [3]int{tempSlice[0], tempSlice[1], tempSlice[2]} // hashMap accepts only arrays(not slices) as key

				if _, exist := hashMapTriplets[hashSlice]; !exist {
					// Hurray, we have not got such triplet earlier
					res = append(res, tempSlice)
					hashMapTriplets[hashSlice] = isMember
				}

			} else {
				// today we dont know if proposal exist, so just mark number as verified
				twoSumVerifiedMap[nums[j]] = isMember
			}
		}
	}

	return res
}

// threeSumFast pre-sort slice to apply 2-pointer 2sum solution and avoid additional loop and hashSet
func threeSumFast(nums []int) [][]int {
	res := [][]int{} // final result

	if len(nums) < 3 {
		return res
	}

	if len(nums) == 3 {
		if nums[0]+nums[1]+nums[2] == 0 {
			return append(res, nums)
		}
		return res
	}

	sort.Ints(nums)                             // sort slice before all operation to apply 2-pointer solution
	var isMember hashSet                        // TODO find a way to check uniq w/ hashSet for sorted slice
	hashMapTriplets := make(map[[3]int]hashSet) // hash map for keeping only unique triplets!
	// A + B + C = 0 is our current expression
	for i := 0; i < len(nums)-1; i++ { // limit range because of nested loop for sum2

		twoNumExpressionTarget := nums[i] // declare A, now we have only B and C as typical 2sum problem
		head := i + 1                     // declare possible B
		tail := len(nums) - 1             // declare possible C

		for head < tail {
			totalSum := nums[head] + nums[tail] + twoNumExpressionTarget // A + B + C

			if totalSum == 0 {
				hashSlice := [3]int{twoNumExpressionTarget, nums[head], nums[tail]} // hashMap accepts only arrays(not slices) as key

				if _, exist := hashMapTriplets[hashSlice]; !exist {
					// Hurray, we have not got such triplet earlier
					res = append(res, []int{twoNumExpressionTarget, nums[head], nums[tail]})
					hashMapTriplets[hashSlice] = isMember
				}
				head++ // move pointer futher to check other combinations
				tail-- // also move, coz there is no chance for a suitable combination
			}

			if totalSum > 0 {
				tail--
			}

			if totalSum < 0 {
				head++
			}

		}
	}

	return res
}

func main() {
	fmt.Println("vim-go")
}
