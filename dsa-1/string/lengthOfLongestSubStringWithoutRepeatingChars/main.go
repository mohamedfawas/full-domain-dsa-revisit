package main

import "fmt"

func lengthOfLongestSubstring(s string) int {
	// Map to keep track of the last index of each character
	lastSeen := make(map[rune]int)
	maxLength := 0
	start := 0 // Start of the sliding window

	for end, currentChar := range s {
		// If the character is seen again and is within the current window
		if lastIndex, exists := lastSeen[currentChar]; exists && lastIndex >= start {
			// Move the start to the right of the previous occurrence
			start = lastIndex + 1
		}

		// Update the last seen index of the current character
		lastSeen[currentChar] = end

		// Calculate current window length and update maxLength if needed
		currentWindowLength := end - start + 1
		if currentWindowLength > maxLength {
			maxLength = currentWindowLength
		}
	}

	return maxLength
}

func main() {
	testCases := []string{
		"abcabcbb",
		"bbbbb",
		"pwwkew",
		"",
		" ",
	}

	for _, s := range testCases {
		fmt.Printf("Input: %q\n", s)
		fmt.Printf("Length of longest substring: %d\n\n", lengthOfLongestSubstring(s))
	}
}
