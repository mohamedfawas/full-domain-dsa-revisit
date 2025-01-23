package main

func FindLongestIncreasingSubArray(arr []int) []int {
	// If array is empty or has only one element, return it as is
	if len(arr) <= 1 {
		return arr
	}

	// Variables to keep track of current sequence
	currentStart := 0
	currentLength := 1

	// Variables to keep track of longest sequence found so far
	maxStart := 0
	maxLength := 1

	// Iterate through the array starting from second element
	for i := 1; i < len(arr); i++ {
		if arr[i] > arr[i-1] {
			// Increment current sequence length
			currentLength++

			// If current sequence is longer than max sequence found so far
			if currentLength > maxLength {
				maxLength = currentLength
				maxStart = currentStart
			}
		} else {
			// Reset current sequence
			currentLength = 1
			currentStart = i
		}
	}

	return arr[maxStart : maxStart+maxLength]
}
