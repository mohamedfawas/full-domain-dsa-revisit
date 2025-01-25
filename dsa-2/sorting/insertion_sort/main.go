package main

import "fmt"

func InsertionSort(arr []int) {
	// Step 1: Get the length of the array
	n := len(arr)

	// Step 2: Start iterating over the array from the second element (index 1)
	// We assume the first element is already "sorted"
	for i := 1; i < n; i++ {
		// Initialize `j` to the current index `i`
		j := i

		// Step 3: Compare the current element `arr[j]` with the element before it (`arr[j-1]`)
		// Keep moving it back until it's in the correct position or `j` becomes 0
		for j > 0 && arr[j-1] > arr[j] {
			// Swap the elements if the previous one is greater than the current one
			// This ensures that smaller elements "bubble" to the left
			arr[j-1], arr[j] = arr[j], arr[j-1]

			// Decrement `j` to compare the element with the next previous one
			j--
		}
	}

	// The array is now sorted in ascending order
}

func main() {
	arr := []int{12, 11, 13, 5, 6}
	fmt.Println("Original array:", arr)
	InsertionSort(arr)
	fmt.Println("Sorted array:", arr)
}
