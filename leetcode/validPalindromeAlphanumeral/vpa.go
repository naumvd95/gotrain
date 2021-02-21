import (
	"strings"
	"unicode"
)

/*
Given a string s, determine if it is a palindrome, considering only alphanumeric characters and ignoring cases.

Example 1:

Input: s = "A man, a plan, a canal: Panama"
Output: true
Explanation: "amanaplanacanalpanama" is a palindrome.

Example 2:

Input: s = "race a car"
Output: false
Explanation: "raceacar" is not a palindrome.

Example 2:

Input: s = "0P"
Output: false

Constraints:

1 <= s.length <= 2 * 105
s consists only of printable ASCII characters.
*/

// time: O(n) | space: O(1)
func isPalindrome(s string) bool {
	l := 0
	r := len(s) - 1

	for l < r {
		// get rid of spec. symbols
		if !isAlphanumeric(s[l]) {
			l++
			continue
		}
		if !isAlphanumeric(s[r]) {
			r--
			continue
		}

		// compare
		lowL := strings.ToLower(string(s[l]))
		lowR := strings.ToLower(string(s[r]))
		if lowL != lowR {
			return false
		}

		// continue loop
		l++
		r--
	}

	return true
}

func isAlphanumeric(c byte) bool {
	r := rune(c)
	if !unicode.IsLetter(r) && !unicode.IsNumber(r) {
		return false
	}

	return true
}
