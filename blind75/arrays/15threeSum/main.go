package main

import "sort"

func threeSum(nums []int) [][]int {
	// Initialize result slice to store all valid triplets
	result := [][]int{}

	// If input array has less than 3 elements, return empty result
	// as we can't form any triplet
	if len(nums) < 3 {
		return result
	}

	// Sort the array first - this is crucial for our two-pointer approach
	// and for handling duplicates efficiently
	// Example: [-4, -1, -1, 0, 1, 2] (sorted array)
	sort.Ints(nums)

	// Iterate through the array up to the third-to-last element
	// as we need at least two more elements after i
	for i := 0; i < len(nums)-2; i++ {
		// Skip duplicates for i to avoid duplicate triplets
		// Example: If we have [-1, -1, -1, 2], we only want to use the first -1
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		// Use two pointers technique: left pointer starts after i
		// right pointer starts at the end of array
		left := i + 1
		right := len(nums) - 1

		// Keep moving pointers while they haven't crossed
		for left < right {
			// Calculate current sum of three numbers
			currentSum := nums[i] + nums[left] + nums[right]

			// If sum is zero, we found a valid triplet
			if currentSum == 0 {
				// Add the triplet to our result
				result = append(result, []int{nums[i], nums[left], nums[right]})

				// Skip duplicates for left pointer
				// Example: If we have [-2, 1, 1, 1], we only want to use the first 1
				for left < right && nums[left] == nums[left+1] {
					left++
				}

				// Skip duplicates for right pointer
				// Example: If we have [-2, 1, 2, 2], we only want to use the last 2
				for left < right && nums[right] == nums[right-1] {
					right--
				}

				// Move both pointers after finding a valid triplet
				left++
				right--
			} else if currentSum < 0 {
				// If sum is less than zero, we need a bigger number
				// So move left pointer to the right
				left++
			} else {
				// If sum is greater than zero, we need a smaller number
				// So move right pointer to the left
				right--
			}
		}
	}

	return result
}

/*
Here's how the algorithm works with example: nums = [-1, 0, 1, 2, -1, -4]

1. First, sort the array: [-4, -1, -1, 0, 1, 2]

2. Start with i = 0 (nums[i] = -4):
   - left = 1, right = 5
   - Try combinations with -4 as first number
   - Move pointers based on sum being less than or greater than 0

3. Move to i = 1 (nums[i] = -1):
   - Skip second -1 to avoid duplicates
   - Find triplets starting with -1
   - When sum is 0 (like -1, 0, 1), add to result

4. Continue this process...

Time Complexity: O(n²) where n is the length of input array
Space Complexity: O(1) excluding the space needed for output

Example Test Cases:
1. [-1,0,1,2,-1,-4] → [[-1,-1,2],[-1,0,1]]
2. [0,0,0] → [[0,0,0]]
3. [1,2,-2,-1] → []
*/
