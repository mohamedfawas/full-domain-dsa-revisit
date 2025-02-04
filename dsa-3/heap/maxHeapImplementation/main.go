package main

import (
	"fmt"
)

// MaxHeap struct represents the max heap data structure.
type MaxHeap struct {
	array []int // Slice to store heap elements
}

// Insert adds a new element to the heap and maintains the heap property.
func (h *MaxHeap) Insert(value int) {
	// Step 1: Add the new element to the end of the array.
	h.array = append(h.array, value)

	// Step 2: Restore the heap property by "bubbling up" the new element.
	h.heapifyUp(len(h.array) - 1)
}

// Extract removes and returns the maximum element (root) from the heap.
func (h *MaxHeap) Extract() (int, bool) {
	// Step 1: Check if the heap is empty.
	if len(h.array) == 0 {
		return -1, false // Return -1 and false if the heap is empty.
	}

	// Step 2: Get the root (maximum element).
	max := h.array[0]

	// Step 3: Replace the root with the last element in the array.
	lastIndex := len(h.array) - 1
	h.array[0] = h.array[lastIndex]
	h.array = h.array[:lastIndex] // Remove the last element.

	// Step 4: Restore the heap property by "bubbling down" the new root.
	h.heapifyDown(0)

	return max, true
}

// heapifyUp restores the heap property by moving an element up the tree.
func (h *MaxHeap) heapifyUp(index int) {
	// Loop until the element is in the correct position.
	for index > 0 {
		// Find the parent index.
		parentIndex := (index - 1) / 2

		// If the current element is less than or equal to its parent, stop.
		if h.array[index] <= h.array[parentIndex] {
			break
		}

		// Swap the current element with its parent.
		h.array[index], h.array[parentIndex] = h.array[parentIndex], h.array[index]

		// Move up to the parent index for the next iteration.
		index = parentIndex
	}
}

// heapifyDown restores the heap property by moving an element down the tree.
func (h *MaxHeap) heapifyDown(index int) {
	// Loop until the element is in the correct position.
	for {
		leftChildIndex := 2*index + 1  // Calculate the left child index.
		rightChildIndex := 2*index + 2 // Calculate the right child index.
		largest := index               // Assume the current index is the largest.

		// Compare the current element with its left child.
		if leftChildIndex < len(h.array) && h.array[leftChildIndex] > h.array[largest] {
			largest = leftChildIndex
		}

		// Compare the current element with its right child.
		if rightChildIndex < len(h.array) && h.array[rightChildIndex] > h.array[largest] {
			largest = rightChildIndex
		}

		// If the largest element is the current element, stop.
		if largest == index {
			break
		}

		// Swap the current element with the largest child.
		h.array[index], h.array[largest] = h.array[largest], h.array[index]

		// Move down to the largest child index for the next iteration.
		index = largest
	}
}

// Example usage of the MaxHeap.
func main() {
	heap := &MaxHeap{}

	// Insert elements into the heap.
	heap.Insert(10)
	heap.Insert(5)
	heap.Insert(20)
	heap.Insert(1)
	fmt.Println("Heap after inserts:", heap.array) // Output: [20 10 5 1]

	// Extract the maximum element.
	max, ok := heap.Extract()
	if ok {
		fmt.Println("Extracted max:", max)             // Output: 20
		fmt.Println("Heap after extract:", heap.array) // Output: [10 1 5]
	}
}
