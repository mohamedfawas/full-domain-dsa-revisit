package main

func maxProduct(nums []int) int {
	// Edge case: If array is empty, return 0
	if len(nums) == 0 {
		return 0
	}

	// Initialize three variables:
	// maxProduct: stores the overall maximum product found so far
	// currMax: keeps track of the maximum product ending at current position
	// currMin: keeps track of the minimum product ending at current position
	// We need to track minimum because when multiplying with a negative number,
	// the minimum can become the maximum
	maxProduct := nums[0]
	currMax := nums[0]
	currMin := nums[0]

	// Example: For array [2,3,-2,4]
	// At index 0: maxProduct = 2, currMax = 2, currMin = 2
	// At index 1: maxProduct = 6, currMax = 6, currMin = 3
	// At index 2: maxProduct = 6, currMax = -2, currMin = -12
	// At index 3: maxProduct = 6, currMax = 4, currMin = -48

	// Iterate through the array starting from index 1
	for i := 1; i < len(nums); i++ {
		// Store current value for easier reference
		curr := nums[i]

		// When multiplying with current number, we have three possibilities:
		// 1. curr * currMax (positive * positive or negative * negative)
		// 2. curr * currMin (might be larger if curr is negative)
		// 3. curr itself (start new subarray from current position)

		// Calculate all possible products
		temp := currMax // Store currMax temporarily as we'll need it for currMin calculation
		currMax = max(max(curr*temp, curr*currMin), curr)
		currMin = min(min(curr*temp, curr*currMin), curr)

		// Update maxProduct if we found a larger product
		maxProduct = max(maxProduct, currMax)
	}

	return maxProduct
}

// Helper function to find maximum of two integers
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// Helper function to find minimum of two integers
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
