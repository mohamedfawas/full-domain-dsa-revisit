package main

import "fmt"

// MergeSort is a function that sorts an array using the merge sort algorithm
func MergeSort(arr []int) []int {
	n := len(arr)
	// Base case: if the array has 1 or 0 elements, it is already sorted
	if n <= 1 {
		return arr
	}

	// Step 1: find the middle index
	mid := n / 2

	// Step 2: recursively split and sort the left half
	left := MergeSort(arr[:mid])

	// Step 3: recursively split and sort the right half
	right := MergeSort(arr[mid:])

	// Step 4: merge the sorted halves and return the result
	return Merge(left, right)
}

// Merge is a helper function that merges two sorted slices into one sorted slice
func Merge(left, right []int) []int {
	// Create a slice to hold the merged result
	result := []int{}

	// Indices for iterating over the left and right slices
	i, j := 0, 0

	// Loop until we have exhausted one of the slices
	for i < len(left) && j < len(right) {
		// Compare elements from both slices and append the smaller one to the result
		if left[i] < right[j] {
			result = append(result, left[i])
			i++
		} else {
			result = append(result, right[j])
			j++
		}
	}

	// Append any remaining elements from the left slice
	for i < len(left) {
		result = append(result, left[i])
		i++
	}

	// Append any remaining elements from the right slice
	for j < len(right) {
		result = append(result, right[j])
		j++
	}

	// Return the merged result
	return result
}

func main() {
	// Example usage of MergeSort
	arr := []int{38, 27, 43, 3, 9, 82, 10}
	fmt.Println("Unsorted array:", arr)
	sortedArr := MergeSort(arr)
	fmt.Println("Sorted array:", sortedArr)
}
