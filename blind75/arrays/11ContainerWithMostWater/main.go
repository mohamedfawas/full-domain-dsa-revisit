package main

func maxArea(height []int) int {
	// Initialize two pointers: left starts at the beginning (0) and right starts at the end (len-1)
	// We'll use these pointers to keep track of the two lines that form our container
	left := 0
	right := len(height) - 1

	// Initialize maxWater to store the maximum area we've found so far
	maxWater := 0

	// Continue until the pointers meet
	// We use two pointers approach because:
	// 1. We need to find two lines to form a container
	// 2. The width of container will be the distance between these lines
	// 3. The height will be limited by the shorter line
	for left < right {
		// Calculate current width between the lines
		// Example: if left = 0 and right = 8, width = 8 - 0 = 8
		width := right - left

		// Calculate the height of the container
		// The height is limited by the shorter line
		// We use min() because water will spill over the shorter line
		// Example: if height[left] = 1 and height[right] = 7, height = 1
		var currentHeight int
		if height[left] < height[right] {
			currentHeight = height[left]
		} else {
			currentHeight = height[right]
		}

		// Calculate current area (width * height)
		// Example: if width = 8 and height = 1, area = 8 * 1 = 8
		currentArea := width * currentHeight

		// Update maxWater if current area is larger
		// This keeps track of the largest container we've found
		if currentArea > maxWater {
			maxWater = currentArea
		}

		// Move the pointer that points to the shorter line
		// We do this because:
		// 1. Area is limited by the shorter line
		// 2. Moving the pointer of taller line will only decrease width
		//    and can't increase height (still limited by shorter line)
		// 3. Moving the shorter line gives us a chance to find a taller line
		//    that might give us a larger area
		if height[left] < height[right] {
			left++
		} else {
			right--
		}
	}

	return maxWater
}

/*
Let's walk through Example 1: height = [1,8,6,2,5,4,8,3,7]

Initial state:
- left = 0 (height = 1)
- right = 8 (height = 7)
- width = 8
- currentHeight = min(1, 7) = 1
- area = 8 * 1 = 8
- maxWater = 8

After few iterations, we reach:
- left = 1 (height = 8)
- right = 8 (height = 7)
- width = 7
- currentHeight = min(8, 7) = 7
- area = 7 * 7 = 49
- maxWater = 49 (This turns out to be our answer)

The algorithm continues checking other possibilities but won't find a larger area.
This is because:
1. To get a larger area, we need either:
   - A larger width (impossible as we started with maximum width)
   - Both lines being taller (impossible as we're limited by the shorter line)
2. As we move pointers inward, width decreases
3. To compensate for decreased width, we'd need much taller lines

Time Complexity: O(n) - we traverse the array once
Space Complexity: O(1) - we only use a constant amount of extra space
*/
