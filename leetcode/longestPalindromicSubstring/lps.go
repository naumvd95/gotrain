package main

/*
Given a string s, return the longest palindromic substring in s.

Example 1:

Input: s = "babad"
Output: "bab"
Note: "aba" is also a valid answer.
Example 2:

Input: s = "cbbd"
Output: "bb"
Example 3:

Input: s = "a"
Output: "a"
Example 4:

Input: s = "ac"
Output: "a"

Constraints:

1 <= s.length <= 1000
s consist of only digits and English letters (lower-case and/or upper-case)
--------------------------------------------------------------------------
Solution idea: Propose palindrome centerS -> compare lengths -> detect left/right shift from centerS to identify palindrome position


1. You dont need to iterate over all possible substrings, identify possible palindrome center and expand left/right sides: check if they are equal

2. What can be palindrome [CENTER]?
 - each letter, except first and last element in array, like "aBa"
 - in between 2 letters, like "aBBA ->> aB|Ba"
 overall it can be n * n - 1 -> n letters, and (n-1) positions between letters

 3. [Iterate] over slice, propose palindrome centers (single letter and in-between) and expand letters around it:
 - 1st center: current position
 - 2nd center: current position and current postion + 1 ->>> in-between center


 4. [Compare] length of generated palindromes with the longest

 5. If new longest detected, determine its position:
 - detect [left shift] from the center (add "-1")
 - detect [right shift] from the center = len of palindrome / 2
--------------------------------------------------------------------------
Complexity
time: O(n^2) ->> n (loop overall) * n(worst case expandAroundCenter)
space: O(1), because we can use the same objects each time loop iteration

- singleLetterCenter int
- inBetweenLetterCenter int
- currentMaxP int
*/

func longestPalindrome(s string) string {
	startLongest := 0
	endLongest := 0
	if len(s) < 1 {
		return ""
	}

	for i := 0; i < len(s); i++ { // i = 0 to handle corner cases
		singleLetterCenter := expandAroundCenter(s, i, i)      // each letter case (odd)
		inBetweenLetterCenter := expandAroundCenter(s, i, i+1) // to cover in-between case (even) we shift to the right each time [REMEMBER +1 shift here]*
		currentMaxP := maxInt(singleLetterCenter, inBetweenLetterCenter)

		if currentMaxP > endLongest-startLongest+1 { // +1 to count len for index 0 as well
			// calculate how long is shift from the palindrome center
			leftShift := (currentMaxP - 1) / 2 // now we need to shift back, because of [REMEMBER +1 shift here]*
			rightShift := currentMaxP / 2

			startLongest = i - leftShift
			endLongest = i + rightShift
		}

	}

	return s[startLongest : endLongest+1] // +1 to cover corner cases like {0,0}
}

func expandAroundCenter(s string, left, right int) int {
	for left >= 0 && right < len(s) { // avoid out-of-range
		if s[left] != s[right] { // while letters mirrored fine
			break
		}

		left--
		right++
	}

	return right - left - 1 // -1 because left or right might be over the range because of ">=" conditions
}

func maxInt(p1, p2 int) int {
	if p1 >= p2 {
		return p1
	}

	return p2
}
