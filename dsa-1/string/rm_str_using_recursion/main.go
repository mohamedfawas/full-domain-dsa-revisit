// Remove a char from a string using recursion.

package main

import "fmt"

func removeChar(s string, target byte) string {
	// Base case: empty string
	if len(s) == 0 {
		return ""
	}

	// Check first character
	if s[0] == target {
		// Skip the character and process the rest
		return removeChar(s[1:], target)
	}
	// Keep the character and process the rest
	return string(s[0]) + removeChar(s[1:], target)
}

func main() {
	original := "hello world"
	target := 'l'
	result := removeChar(original, byte(target))

	fmt.Printf("Original: %s\n", original)
	fmt.Printf("After removing '%c': %s\n", target, result)
}
