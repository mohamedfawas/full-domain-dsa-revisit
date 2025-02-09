package main

func search(nums []int, target int) int {
	// If array is empty, return -1
	if len(nums) == 0 {
		return -1
	}

	// Initialize two pointers for binary search
	left := 0
	right := len(nums) - 1

	// Continue binary search while left pointer is less than or equal to right pointer
	for left <= right {
		// Calculate middle index
		// We use (left + right) / 2 to find the middle
		mid := (left + right) / 2

		// If we found the target at middle position, return its index
		if nums[mid] == target {
			return mid
		}

		// Now we need to determine which half of the array is sorted
		// This is important because one half will always be sorted in a rotated sorted array

		// Check if the left half is sorted
		if nums[left] <= nums[mid] {
			// Example: [4,5,6,7,0,1,2], left = 0, mid = 3, right = 6
			// Left half is [4,5,6,7] which is sorted

			// Check if target lies in the left half
			// Target should be greater than or equal to leftmost element
			// AND less than middle element
			if target >= nums[left] && target < nums[mid] {
				// Target is in left half, so search there
				right = mid - 1
			} else {
				// Target is in right half
				left = mid + 1
			}
		} else {
			// Right half is sorted
			// Example: [6,7,0,1,2,4,5], left = 0, mid = 3, right = 6
			// Right half is [1,2,4,5] which is sorted

			// Check if target lies in the right half
			// Target should be greater than middle element
			// AND less than or equal to rightmost element
			if target > nums[mid] && target <= nums[right] {
				// Target is in right half, so search there
				left = mid + 1
			} else {
				// Target is in left half
				right = mid - 1
			}
		}
	}

	// If we reach here, target was not found
	return -1
}
