package main

func findMin(nums []int) int {
	// If array has only one element, return that element
	// Example: nums = [5] -> return 5
	if len(nums) == 1 {
		return nums[0]
	}

	// Initialize two pointers for binary search
	// left points to start of array, right points to end
	left := 0
	right := len(nums) - 1

	// If array is not rotated (i.e., already sorted)
	// Example: nums = [1,2,3,4,5] -> return nums[0] which is 1
	if nums[right] > nums[left] {
		return nums[left]
	}

	// Binary search implementation
	// We keep searching until left and right pointers meet
	for left < right {
		// Calculate middle index
		// Example: If left = 0 and right = 4, mid = 2
		mid := left + (right-left)/2

		// If middle element is greater than next element
		// we found the pivot point (minimum element)
		// Example: nums = [3,4,5,1,2], when mid = 2:
		// nums[2] = 5 > nums[3] = 1, so return 1
		if nums[mid] > nums[mid+1] {
			return nums[mid+1]
		}

		// If middle element is less than previous element
		// we found the pivot point (current element is minimum)
		// Example: nums = [4,5,1,2,3], when mid = 2:
		// nums[1] = 5 > nums[2] = 1, so return 1
		if nums[mid-1] > nums[mid] {
			return nums[mid]
		}

		// If middle element is greater than first element
		// minimum lies in right half
		// Example: nums = [4,5,6,1,2], when mid = 2:
		// nums[2] = 6 > nums[0] = 4, so search right half
		if nums[mid] > nums[left] {
			left = mid + 1
		} else {
			// Otherwise, minimum lies in left half
			// Example: nums = [5,1,2,3,4], when mid = 2:
			// nums[2] = 2 < nums[0] = 5, so search left half
			right = mid - 1
		}
	}

	// This line should never be reached if input follows constraints
	// but we return first element as a safeguard
	return nums[0]
}
