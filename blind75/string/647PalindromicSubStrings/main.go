package main

// countSubstrings counts the total number of palindromic substrings in the given string
// A palindrome is a string that reads the same forwards and backwards
// For example: in string "aaa", palindromes are "a", "a", "a", "aa", "aa", "aaa"
func countSubstrings(s string) int {
	// Initialize total count of palindromes
	// Every single character is a palindrome by itself,
	// so we'll find more as we expand around centers
	totalCount := 0

	// Helper function to count palindromes by expanding around a center
	// It takes left and right indices and expands outward while characters match
	expandAroundCenter := func(left, right int) {
		// Keep expanding as long as:
		// 1. left index doesn't go below 0
		// 2. right index doesn't exceed string length
		// 3. characters at both positions match
		for left >= 0 && right < len(s) && s[left] == s[right] {
			// We found a palindrome! Increment our count
			totalCount++

			// Move outward in both directions to check for larger palindromes
			left--
			right++
		}
	}

	// Iterate through each character as potential center
	for i := 0; i < len(s); i++ {
		// For each position, we need to check two types of palindromes:

		// 1. Odd length palindromes (single character center)
		// Example: For "aba", center is 'b'
		expandAroundCenter(i, i)

		// 2. Even length palindromes (between two characters)
		// Example: For "aa", center is between both 'a's
		expandAroundCenter(i, i+1)
	}

	return totalCount
}

// Let's see how it works with examples:
/*
Example 1: s = "abc"
- For position 0 ('a'):
  * Odd length: "a" (count = 1)
  * Even length: no palindrome with 'a' and 'b'
- For position 1 ('b'):
  * Odd length: "b" (count = 2)
  * Even length: no palindrome with 'b' and 'c'
- For position 2 ('c'):
  * Odd length: "c" (count = 3)
  * Even length: no palindrome (out of bounds)
Total count = 3

Example 2: s = "aaa"
- For position 0 (first 'a'):
  * Odd length: "a" (count = 1)
  * Even length: "aa" (count = 2)
- For position 1 (second 'a'):
  * Odd length: "a" (count = 3)
  * Even length: "aa" (count = 4)
- For position 2 (third 'a'):
  * Odd length: "a" (count = 5)
  * Even length: no palindrome (out of bounds)
- Additionally, "aaa" is found when expanding from center position 1
  (count = 6)
Total count = 6
*/

// Step by step visualization for "aaa":
// 1. Start with empty count = 0
// 2. At index 0:
//    - Check "a" (odd length) ✓ count = 1
//    - Check "aa" (even length) ✓ count = 2
// 3. At index 1:
//    - Check "a" (odd length) ✓ count = 3
//    - Check "aa" (even length) ✓ count = 4
//    - While expanding odd length, find "aaa" ✓ count = 5
// 4. At index 2:
//    - Check "a" (odd length) ✓ count = 6
//    - Check even length (no valid palindrome)
