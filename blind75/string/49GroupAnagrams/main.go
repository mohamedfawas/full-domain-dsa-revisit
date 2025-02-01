package main

import (
	"sort"
	"strings"
)

// This function takes a slice of strings and groups anagrams together
// For example: Input: ["eat", "tea", "tan", "ate", "nat", "bat"]
// Output: [["eat","tea","ate"], ["tan","nat"], ["bat"]]
func groupAnagrams(strs []string) [][]string {
	// Create a map where:
	// - The key is the sorted version of a word (e.g., "aet" for both "eat" and "tea")
	// - The value is a slice containing all words that are anagrams of each other
	anagramMap := make(map[string][]string)

	// Loop through each string in the input slice
	// For example, if str is "eat":
	for _, str := range strs {
		// Split the string into a slice of characters
		// "eat" becomes ["e", "a", "t"]
		chars := strings.Split(str, "")

		// Sort the characters alphabetically
		// ["e", "a", "t"] becomes ["a", "e", "t"]
		sort.Strings(chars)

		// Join the sorted characters back into a string
		// ["a", "e", "t"] becomes "aet"
		sortedStr := strings.Join(chars, "")

		// Add the original string to the map using the sorted string as the key
		// If str is "eat", it will be stored under key "aet"
		// If later we find "tea", it will also be stored under "aet"
		anagramMap[sortedStr] = append(anagramMap[sortedStr], str)
	}

	// Create a slice to store the final result
	var result [][]string

	// Loop through each group of anagrams in the map
	// For example, anagramMap["aet"] might contain ["eat", "tea", "ate"]
	for _, group := range anagramMap {
		// Add each group of anagrams to the result slice
		result = append(result, group)
	}

	// Return the grouped anagrams
	return result
}
