package main

/*
Given an array a that contains only numbers in the range from 1 to a.length, find the first duplicate number for which the second occurrence has the minimal index. In other words, if there are more than 1 duplicated numbers, return the number for which the second occurrence has a smaller index than the second occurrence of the other number does. If there are no such elements, return -1.

*/

// Golang implementation of set is map[SMTH]bool or map[SMTH]struct{}, second is more lightweight
type hashSet struct{}

func GetFirstDuplicate(data []int) int {
	var isMember hashSet
	set := make(map[int]hashSet)

	for _, v := range data {
		if _, ok := set[v]; ok {
			return v
		}
		set[v] = isMember
	}
	return -1
}
