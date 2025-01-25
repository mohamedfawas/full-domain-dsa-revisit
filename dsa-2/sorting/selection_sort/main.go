package main

import "fmt"

// SelectionSort function sorts an array using selection sort algorithm
func SelectionSort(arr []int) {
	n := len(arr)
	for i := 0; i < n-1; i++ {
		// Assume the minimum is the first element
		minIndex := i
		// Iterate through the unsorted elements
		for j := i + 1; j < n; j++ {
			// If the current element is smaller than the minimum, update minIndex
			if arr[j] < arr[minIndex] {
				minIndex = j
			}
		}
		// Swap the found minimum element with the first element
		if minIndex != i {
			arr[i], arr[minIndex] = arr[minIndex], arr[i]
		}
	}
}

func main() {
	arr := []int{64, 25, 12, 22, 11}
	fmt.Println("Unsorted array:", arr)
	SelectionSort(arr)
	fmt.Println("Sorted array:", arr)
}
