package main

func longestSubStringWithoutRepeatingChars(s string) string {
	lastSeen := make(map[rune]int)

	start := 0
	maxLength := 0
	maxStart := 0

	for end, currentChar := range s {
		if lastIndex, exists := lastSeen[currentChar]; exists && lastIndex >= start {
			start = lastIndex + 1
		}

		lastSeen[currentChar] = end

		currentLength := end - start + 1
		if currentLength > maxLength {
			maxLength = currentLength
			maxStart = start
		}
	}

	return s[maxStart : maxStart+maxLength]
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
		println(longestSubStringWithoutRepeatingChars(s))
	}
}
