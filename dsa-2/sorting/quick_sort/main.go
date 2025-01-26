package main

// partition rearranges the elements in the array such that elements less than the pivot are on the left,
// and elements greater than the pivot are on the right. It returns the index of the pivot element after partitioning.
// arr: the array to be partitioned
// lb: the lower bound index of the array segment to be partitioned
// ub: the upper bound index of the array segment to be partitioned
func partition(arr []int, lb, ub int) int {
	// Choose the pivot element from the lower bound of the array segment
	pivot := arr[lb]
	start := lb
	end := ub
	// Loop until the start index is less than the end index
	for start < end {
		// Increment start index until an element greater than the pivot is found
		for start < ub && arr[start] <= pivot {
			start++
		}
		// Decrement end index until an element less than or equal to the pivot is found
		for arr[end] > pivot {
			end--
		}
		// Swap the elements at start and end indices if start is less than end
		if start < end {
			arr[start], arr[end] = arr[end], arr[start]
		}
	}

	// Swap the pivot element with the element at the end index
	arr[end], arr[lb] = arr[lb], arr[end]
	// Return the index of the pivot element after partitioning
	return end
}

// QuickSort sorts the array using the quicksort algorithm.
// arr: the array to be sorted
// lb: the lower bound index of the array segment to be sorted
// ub: the upper bound index of the array segment to be sorted
func QuickSort(arr []int, lb, ub int) {
	// Check if the lower bound is less than the upper bound
	if lb < ub {
		// Partition the array and get the pivot index
		loc := partition(arr, lb, ub)
		// Recursively sort the elements before the pivot
		QuickSort(arr, lb, loc-1)
		// Recursively sort the elements after the pivot
		QuickSort(arr, loc+1, ub)
	}
}
