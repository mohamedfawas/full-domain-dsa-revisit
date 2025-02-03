package main

// This problem can be solved using Kadane's Algorithm
// The idea is to keep track of:
// 1. The maximum sum ending at current position (currentSum)
// 2. The overall maximum sum found so far (maxSum)

func maxSubArray(nums []int) int {
	// Edge case: if array has only one element, return that element
	if len(nums) == 1 {
		return nums[0]
	}

	// Initialize currentSum and maxSum with the first element
	currentSum := nums[0]
	maxSum := nums[0]

	// Start from second element (index 1)
	// For each element, we have two choices:
	// 1. Add it to the existing subarray (currentSum + nums[i])
	// 2. Start a new subarray from this element (nums[i])
	for i := 1; i < len(nums); i++ {
		// If currentSum becomes negative, it's better to start fresh
		// Example: if currentSum is -2 and next number is 4
		// It's better to start new subarray from 4 rather than having -2 + 4 = 2
		if currentSum < 0 {
			currentSum = nums[i]
		} else {
			// If currentSum is positive, add current element to it
			currentSum = currentSum + nums[i]
		}

		// Update maxSum if currentSum becomes larger
		// This keeps track of the largest sum we've seen so far
		if currentSum > maxSum {
			maxSum = currentSum
		}
	}

	return maxSum
}

/* Let's see how this works with Example 1: nums = [-2,1,-3,4,-1,2,1,-5,4]

Step by step:
1. Initialize: currentSum = -2, maxSum = -2
2. For 1:
   - currentSum < 0, so currentSum = 1
   - maxSum = 1 (as 1 > -2)
3. For -3:
   - currentSum = 1 + (-3) = -2
   - maxSum still 1
4. For 4:
   - currentSum < 0, so currentSum = 4
   - maxSum = 4 (as 4 > 1)
5. For -1:
   - currentSum = 4 + (-1) = 3
   - maxSum still 4
6. For 2:
   - currentSum = 3 + 2 = 5
   - maxSum = 5
7. For 1:
   - currentSum = 5 + 1 = 6
   - maxSum = 6
8. For -5:
   - currentSum = 6 + (-5) = 1
   - maxSum still 6
9. For 4:
   - currentSum = 1 + 4 = 5
   - maxSum still 6

Final result: 6 */
