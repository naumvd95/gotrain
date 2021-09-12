package programmerstring

import "fmt"

// time: O(n) coz there are 2 separate loops l->r and r->l | space: O(2m) where m is a uniq letters in deesired word
func lengthBetweenWords(desiredWord, s string) int {
	// 1. loop from left-> right and find first word and its last letter-idx
	// 2. loop from right->left and find last word and its last letter-idx
	// 2. calculate diff between them
	validWordTemplate := make(map[byte]int) // first space O(m)
	for i := 0; i < len(desiredWord); i++ {
		letter := desiredWord[i]
		validWordTemplate[letter]++
	}
	fmt.Printf("Hey, word template is: %v\n", validWordTemplate)

	// 1. find first occurence of desiredWord
	wordTemplate := make(map[byte]int)   // second space O(m)
	for key := range validWordTemplate { // time O(m)
		wordTemplate[key] = 0
	}
	start := 0
	for start < len(s) && len(wordTemplate) > 0 { // time O(n)
		currentLetter := s[start]

		if _, letterExists := wordTemplate[currentLetter]; letterExists {
			wordTemplate[currentLetter]++

			// if we found more than enough occurencies for letter -> delete such key
			if wordTemplate[currentLetter] >= validWordTemplate[currentLetter] {
				delete(wordTemplate, currentLetter)
				fmt.Printf(" we found more than enough occurencies for letter %v -> delete  key %v: current state: %v\n", string(currentLetter), currentLetter, wordTemplate)
			}
		}
		start++
	}
	fmt.Printf("first word ends at: %v\n", start)

	// 2. find last occurence of desiredWord
	for key := range validWordTemplate { // time O(m)
		wordTemplate[key] = 0
	}
	end := len(s) - 1
	for end >= 0 && len(wordTemplate) > 0 { // time O(n)
		currentLetter := s[end]

		if _, letterExists := wordTemplate[currentLetter]; letterExists {
			wordTemplate[currentLetter]++

			// if we found more than enough occurencies for letter -> delete such key
			if wordTemplate[currentLetter] >= validWordTemplate[currentLetter] {
				delete(wordTemplate, currentLetter)
				fmt.Printf(" we found more than enough occurencies for letter %v -> delete  key %v : current state: %v\n", string(currentLetter), currentLetter, wordTemplate)
			}
		}
		end--
	}
	fmt.Printf("last word starts at: %v\n", end)

	return end - start + 1
}
