package main

// ReverseStringUsingStack takes a string input and returns its reverse using a stack data structure
// Example: If input is "hello", it returns "olleh"
func ReverseStringUsingStack(input string) string {
	// Create an empty slice of runes to use as our stack
	// We use rune type instead of byte to properly handle Unicode characters
	// Example: For input "hello", stack will eventually contain ['h','e','l','l','o']
	stack := []rune{}

	// Iterate through each character (rune) in the input string
	// The range keyword gives us each Unicode character (rune) from the string
	// Example: If input is "hello", char will be 'h', then 'e', then 'l', etc.
	for _, char := range input {
		// Append each character to the end of our stack
		// This is like pushing elements onto a stack
		// After this loop for "hello", stack = ['h','e','l','l','o']
		stack = append(stack, char)
	}

	// Initialize an empty string to store our reversed result
	reversed := ""

	// Continue loop as long as there are elements in the stack
	// This loop processes: ['h','e','l','l','o'] -> "olleh"
	for len(stack) > 0 {
		// Get the index of the last element in the stack
		// Example: If stack = ['h','e','l','l','o'], then n = 4
		n := len(stack) - 1

		// Add the last character from stack to our reversed string
		// string(stack[n]) converts the rune back to a string
		// Example: If stack = ['h','e','l','l','o'], stack[n] = 'o'
		reversed += string(stack[n])

		// Remove the last element from the stack using slice operations
		// stack[:n] creates a new slice excluding the last element
		// Example: ['h','e','l','l','o'] becomes ['h','e','l','l']
		stack = stack[:n]
	}

	// Return the fully reversed string
	// Example: For input "hello", returns "olleh"
	return reversed
}
