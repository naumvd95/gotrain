package main

/*
Given a string s, find the length of the longest substring without repeating characters.

Example 1:

Input: s = "abcabcbb"
Output: 3
Explanation: The answer is "abc", with the length of 3.
Example 2:

Input: s = "bbbbb"
Output: 1
Explanation: The answer is "b", with the length of 1.
Example 3:

Input: s = "pwwkew"
Output: 3
Explanation: The answer is "wke", with the length of 3.
Notice that the answer must be a substring, "pwke" is a subsequence and not a substring.
Example 4:

Input: s = ""
Output: 0


Constraints:

0 <= s.length <= 5 * 104
s consists of English letters, digits, symbols and spaces.
*/

func lengthOfLongestSubstring(s string) int {
	longest := 0
	verifiedChars := make(map[byte]int) // key: char, value: its index

	// Sliding window solution
	l := 0                        // left Side Window
	for r := 0; r < len(s); r++ { // right Side Window
		//fmt.Printf("we are in Right window loop, r: %v, verified chars: %v\n", r, verifiedChars)

		if _, ok := verifiedChars[s[r]]; ok { // we already met such char

			// shift window futher, right after this char previous appearance
			shiftProposal := verifiedChars[s[r]] + 1
			if shiftProposal > l {
				// only if current left window state less than proposal,
				// otherwise we need to clean verifiedChars from chars, which indicies less than current Left Side Window!
				l = shiftProposal
			}

			//fmt.Printf("we met this var %v before in %v, shifting left futher on position: %v\n", s[r], verifiedChars, l)
		}

		verifiedChars[s[r]] = r // mark char and its position as visited
		//fmt.Printf("we are in Right window loop, updating verified chars: %v\n", verifiedChars)

		counter := r - l + 1 // because index starting from 0
		if counter > longest {
			longest = counter
		}

		//fmt.Printf("we are in Right window loop, longest: %v\n", longest)
	}

	return longest
}
