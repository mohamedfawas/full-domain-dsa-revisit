package main

// productExceptSelf calculates the product of all elements except the current element
func productExceptSelf(nums []int) []int {
	n := len(nums)
	// Initialize result array with 1s
	// We'll use this to store our final answer
	answer := make([]int, n)
	for i := range answer {
		answer[i] = 1
	}

	// Step 1: Calculate prefix products
	// For each position i, we calculate product of all numbers to its left
	// Example with [1,2,3,4]:
	// First pass: answer = [1,1,2,6]
	prefix := 1
	for i := 0; i < n; i++ {
		answer[i] = prefix
		prefix *= nums[i]
	}

	// Step 2: Calculate suffix products and combine with prefix products
	// For each position i, we multiply the prefix product (already in answer[i])
	// with the product of all numbers to its right
	// Example continuation:
	// Second pass: answer = [24,12,8,6]
	suffix := 1
	for i := n - 1; i >= 0; i-- {
		answer[i] *= suffix
		suffix *= nums[i]
	}

	return answer
}

// Example usage:
/*
Let's walk through the example: nums = [1,2,3,4]

First Pass (Prefix products):
1. i=0: prefix=1, answer[0]=1, prefix=1*1=1
2. i=1: prefix=1, answer[1]=1, prefix=1*2=2
3. i=2: prefix=2, answer[2]=2, prefix=2*3=6
4. i=3: prefix=6, answer[3]=6, prefix=6*4=24
After first pass: answer = [1,1,2,6]

Second Pass (Suffix products and combining):
1. i=3: suffix=1, answer[3]=6*1=6, suffix=1*4=4
2. i=2: suffix=4, answer[2]=2*4=8, suffix=4*3=12
3. i=1: suffix=12, answer[1]=1*12=12, suffix=12*2=24
4. i=0: suffix=24, answer[0]=1*24=24, suffix=24*1=24
Final result: answer = [24,12,8,6]

Each position i in the final answer contains:
- Product of all numbers to the left of i (from first pass)
- Multiplied by product of all numbers to the right of i (from second pass)
*/
