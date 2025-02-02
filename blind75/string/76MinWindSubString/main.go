package main

import (
	"fmt"
	"math"
)

func minWindow(s string, t string) string {
	// First, handle the edge case where s is shorter than t
	// Example: s = "ab", t = "abc" -> impossible to find a window containing all chars of t
	if len(s) < len(t) {
		return ""
	}

	// Create two arrays of size 128 to store character counts
	// We use 128 because it covers all ASCII characters
	// These arrays will act as hash tables where the index represents the ASCII value of the character
	targetCounts := make([]int, 128) // Stores how many of each character we need
	windowCounts := make([]int, 128) // Stores how many of each character we currently have

	// Count how many unique characters we need to find
	// Example: for t = "ABC", required = 3
	// Example: for t = "AAB", required = 2 (because A is counted only once as unique)
	required := 0
	for _, char := range t {
		if targetCounts[char] == 0 {
			required++ // First time seeing this character
		}
		targetCounts[char]++ // Increment the count needed for this character
	}

	// Initialize variables for the sliding window
	left, right := 0, 0     // Window boundaries
	formed := 0             // Counts how many unique characters have met their required frequency
	minLen := math.MaxInt32 // Length of the smallest valid window found
	start := 0              // Starting index of the minimum window

	// Example walkthrough with s = "ADOBECODEBANC", t = "ABC"
	// Initial state: left = 0, right = 0, window = ""

	// Main sliding window loop
	for right < len(s) {
		// Get the character at the right pointer and add it to our window
		charRight := s[right]
		windowCounts[charRight]++

		// Check if adding this character completed the required count for it
		// Example: if we need 2 'A's and we just got our second 'A'
		if windowCounts[charRight] == targetCounts[charRight] {
			formed++ // We've satisfied the requirement for one more unique character
		}

		// If we have all required characters, try to minimize the window
		// by moving the left pointer
		for formed == required && left <= right {
			currentWindowLen := right - left + 1

			// If this window is smaller than our previous best, update it
			// Example: if current window "BANC" (length 4) is smaller than "ADOBEC" (length 6)
			if currentWindowLen < minLen {
				minLen = currentWindowLen
				start = left
			}

			// Try removing the leftmost character
			charLeft := s[left]

			// If removing this character would break our required count
			// Example: if we need 1 'A' and we're about to remove our only 'A'
			if windowCounts[charLeft] == targetCounts[charLeft] {
				formed-- // We no longer satisfy this character's requirement
			}

			windowCounts[charLeft]-- // Remove the character from our window
			left++                   // Move left pointer to try a smaller window
		}

		right++ // Move right pointer to try including the next character
	}

	// If we never found a valid window, return empty string
	// This happens when minLen was never updated from its initial value
	if minLen == math.MaxInt32 {
		return ""
	}

	// Return the substring of the minimum window
	// Example: if start = 8 and minLen = 4, return s[8:12] which is "BANC"
	return s[start : start+minLen]
}

func main() {
	// Example 1: Regular case
	// s = "ADOBECODEBANC", t = "ABC"
	// Expected output: "BANC"
	// Explanation: "BANC" contains all characters A, B, and C at least once
	s1 := "ADOBECODEBANC"
	t1 := "ABC"
	fmt.Println(minWindow(s1, t1))

	// Example 2: Single character case
	// s = "a", t = "a"
	// Expected output: "a"
	// Explanation: The entire string is the minimum window
	s2 := "a"
	t2 := "a"
	fmt.Println(minWindow(s2, t2))

	// Example 3: Impossible case
	// s = "a", t = "aa"
	// Expected output: ""
	// Explanation: We need two 'a's but the string only has one
	s3 := "a"
	t3 := "aa"
	fmt.Println(minWindow(s3, t3))
}
