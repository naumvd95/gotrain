package main

import (
	"sort"
	"strings"
)

/*
Given an array of strings strs, group the anagrams together. You can return the answer in any order.

An Anagram is a word or phrase formed by rearranging the letters of a different word or phrase, typically using all the original letters exactly once.



Example 1:

Input: strs = ["eat","tea","tan","ate","nat","bat"]
Output: [["bat"],["nat","tan"],["ate","eat","tea"]]
Example 2:

Input: strs = [""]
Output: [[""]]
Example 3:

Input: strs = ["a"]
Output: [["a"]]

Constraints:

1 <= strs.length <= 104
0 <= strs[i].length <= 100
strs[i] consists of lower-case English letters.
*/

/* groupAnagramsByAlphabetHash will work if we know the strs[i] consists only of lower-case English letters.
COMPLEXITY:
time: O(strMaxLen*strAmount)
space: O(strMaxLen*strAmount) = hash map

SOLUTION MINDSET:
You defenitely need to use hashSet for collecting data "anagram:amount"
1) Use sorted string as hash key
2) OR create alphabet template to apply for each anagram: 0 - means letter is not exist in aragram, 1 and more - means letters exists in anagram
   with such approach, you dont need to sort each anagram string, you just need to split it on chars and apply to your alphabet template, then use it as key in HashSet

*/
func groupAnagramsByAlphabetHash(strs []string) [][]string {
	res := [][]string{}
	anagramMap := make(map[[26]int][]string) //hashMap for detecting groups

	// use alphabet template instead of sorting each string: "aab" string will be represented as []int{2,1,0,0.....
	alphabetHashTemplate := [26]int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
	for _, s := range strs {
		hash := alphabetHashTemplate
		for _, v := range s { // runes
			charPosition := v - 'a' // for example : 'd' - 'a' == 100 - 97 = 3 ->> position in our alphabetHashTemplate
			hash[charPosition]++
		}
		anagramMap[hash] = append(anagramMap[hash], s)
	}

	for _, anagroups := range anagramMap {
		res = append(res, anagroups)
	}

	return res
}

/* groupAnagramsByStringSort will use sorted string as key for hash map
COMPLEXITY:
time: O(strMaxLen*strAmount* log strMaxLen) coz we use sort for strings
space: O(strMaxLen*strAmount) = hash map
*/
func groupAnagramsByStringSort(strs []string) [][]string {
	res := [][]string{}
	anagramMap := make(map[string][]string) //hashMap for detecting groups

	for _, s := range strs {
		tmpS := strings.Split(s, "")
		sort.Strings(tmpS)
		hashString := strings.Join(tmpS, "")

		anagramMap[hashString] = append(anagramMap[hashString], s)
	}

	for _, anagroups := range anagramMap {
		res = append(res, anagroups)
	}

	return res
}
