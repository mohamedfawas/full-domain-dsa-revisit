package main

// longestPalindrome finds the longest palindromic substring in the given string
// A palindrome is a string that reads the same forwards and backwards
// For example: "aba", "bb", "radar" are palindromes
func longestPalindrome(s string) string {
	// If string is empty or has only one character, it's already a palindrome
	if len(s) < 2 {
		return s
	}

	// Variables to keep track of the longest palindrome found
	start := 0     // Starting index of the longest palindrome
	maxLength := 1 // Length of the longest palindrome

	// Helper function to expand around a center
	// It takes left and right indices and expands outward while characters match
	expandAroundCenter := func(left, right int) {
		// Keep expanding as long as we're within string bounds
		// and characters at both ends match
		for left >= 0 && right < len(s) && s[left] == s[right] {
			// Calculate current palindrome length
			currentLength := right - left + 1

			// If this palindrome is longer than our previous longest,
			// update our tracking variables
			if currentLength > maxLength {
				start = left
				maxLength = currentLength
			}

			// Move outward in both directions
			left--
			right++
		}
	}

	// Iterate through each character as potential center
	for i := 0; i < len(s); i++ {
		// Example: For string "babad":
		// When i = 1 ('a'), we check:
		// 1. Odd length palindromes centered at 'a'
		// 2. Even length palindromes centered between 'a' and the next character

		// Check for odd length palindromes (single character center)
		// Example: "aba" - center is 'b'
		expandAroundCenter(i, i)

		// Check for even length palindromes (between two characters)
		// Example: "bb" - center is between both 'b's
		expandAroundCenter(i, i+1)
	}

	// Return the substring using the start index and maxLength
	// Example: if start = 1 and maxLength = 3 for "babad",
	// it returns "bab" (indices 1 to 3)
	return s[start : start+maxLength]
}

// Example usage:
/*
1. s = "babad"
   - First iteration (i = 0, char = 'b'):
     * Odd length: "b" (length 1)
     * Even length: "ba" (not a palindrome)
   - Second iteration (i = 1, char = 'a'):
     * Odd length: "bab" (length 3) ✓ This becomes our current longest
     * Even length: "ab" (not a palindrome)
   - Further iterations don't find longer palindromes
   Result: "bab"

2. s = "cbbd"
   - First iteration (i = 0, char = 'c'):
     * Odd length: "c" (length 1)
     * Even length: "cb" (not a palindrome)
   - Second iteration (i = 1, char = 'b'):
     * Odd length: "b" (length 1)
     * Even length: "bb" (length 2) ✓ This becomes our longest
   Result: "bb"
*/
