package main

// countSubstrings counts all palindromic substrings in a given string
// For example, if s = "aaa":
// - Single characters: "a", "a", "a" (all single characters are palindromes)
// - Two characters: "aa", "aa" (we check positions 0,1 and 1,2)
// - Three characters: "aaa" (we check the whole string)
// Total count would be 6
func countSubstrings(s string) int {
	// Initialize our counter for palindromes
	count := 0

	// For each character in the string, we'll try to expand around it
	// to find all possible palindromes
	for i := 0; i < len(s); i++ {
		// Handle odd-length palindromes
		// Example: for "aba", we start at 'b' and expand outwards
		expandAroundCenter(s, i, i, &count)

		// Handle even-length palindromes
		// Example: for "abba", we start between the two 'b's
		expandAroundCenter(s, i, i+1, &count)
	}

	return count
}

// expandAroundCenter checks if we have a palindrome by expanding around a center point
// Parameters:
// - s: the input string
// - left: left pointer starting position
// - right: right pointer starting position
// - count: pointer to our counter to update when we find palindromes
func expandAroundCenter(s string, left int, right int, count *int) {
	// Keep expanding outwards as long as:
	// 1. We haven't reached the string boundaries
	// 2. Characters at both ends match
	for left >= 0 && right < len(s) && s[left] == s[right] {
		// We found a palindrome! Increment our counter
		*count++

		// Move outwards to check for larger palindromes
		// Example: if we found "aba" is a palindrome,
		// next we check if "xabax" is also a palindrome
		left--
		right++
	}
}

/*
Let's walk through an example: s = "aaa"

1. First iteration (i = 0):
   - Odd-length: starts at 'a' (index 0)
     * "a" is a palindrome (count = 1)
   - Even-length: starts between index 0 and 1
     * "aa" is a palindrome (count = 2)

2. Second iteration (i = 1):
   - Odd-length: starts at 'a' (index 1)
     * "a" is a palindrome (count = 3)
   - Even-length: starts between index 1 and 2
     * "aa" is a palindrome (count = 4)

3. Third iteration (i = 2):
   - Odd-length: starts at 'a' (index 2)
     * "a" is a palindrome (count = 5)
   - Even-length: starts between index 2 and 3
     (this won't find anything as we're at the end)

Note: During these iterations, when we check odd-length palindromes
at index 1, we also find "aaa" as we expand outwards (count = 6)

Final count: 6 palindromes
*/
