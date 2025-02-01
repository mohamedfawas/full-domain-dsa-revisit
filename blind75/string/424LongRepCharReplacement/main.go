package main

func characterReplacement(s string, k int) int {
	// Create an array of size 26 to store frequency of each uppercase letter (A-Z)
	// For example: freq[0] represents count of 'A', freq[1] represents count of 'B', and so on
	// We use 26 because there are 26 uppercase letters in English alphabet
	freq := make([]int, 26)

	// maxWindowSize: keeps track of the longest valid substring we've found so far
	// maxFreq: stores the count of the most frequent character in current window
	// start: left pointer of our sliding window (where our current substring begins)
	maxWindowSize := 0
	maxFreq := 0
	start := 0

	// Loop through the string using 'end' as our right pointer
	// This will help us grow our window from left to right
	for end := 0; end < len(s); end++ {
		// Convert current character to a number between 0-25
		// Example: 'A' becomes 0, 'B' becomes 1, etc.
		// We subtract 'A' to get this conversion:
		// 'A' - 'A' = 0, 'B' - 'A' = 1, and so on
		currentChar := s[end] - 'A'

		// Increment the frequency count for current character
		freq[currentChar]++

		// If we found a new character with higher frequency,
		// update maxFreq to this new highest frequency
		if freq[currentChar] > maxFreq {
			maxFreq = freq[currentChar]
		}

		// Calculate size of current window
		// Example: if start = 1 and end = 3, window size is 3-1+1 = 3 characters
		currentWindowSize := end - start + 1

		// Check if our current window is valid:
		// Window is invalid if: (window size - most frequent char count) > k
		// This means we need more replacements than we're allowed (k)
		// Example: in "ABAB" with k=1:
		// If window="ABAB", size=4, maxFreq=2 (2 A's or 2 B's)
		// 4-2 > 1, so we need more than 1 replacement
		if currentWindowSize-maxFreq > k {
			// Window is too big, need to shrink it from left
			// Decrease frequency count of character at 'start'
			freq[s[start]-'A']--
			// Move start pointer right to shrink window
			start++
		}

		// Check if current window is bigger than our previous maximum
		// If so, update maxWindowSize
		// We don't need to check if window is valid because we've
		// already shrunk any invalid windows in the step above
		if end-start+1 > maxWindowSize {
			maxWindowSize = end - start + 1

		}
	}

	// Return the size of the longest valid window we found
	return maxWindowSize
}
