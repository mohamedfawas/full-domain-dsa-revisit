package main

func maxProfit(prices []int) int {
	if len(prices) <= 1 {
		return 0
	}

	minPrice := prices[0] // Track minimum price seen so far
	maxProfit := 0        // Track maximum profit possible

	// Iterate through prices starting from second day
	for i := 1; i < len(prices); i++ {
		// Update minimum price if current price is lower
		if prices[i] < minPrice {
			minPrice = prices[i]
		} else {
			// Calculate potential profit if we sell at current price
			currentProfit := prices[i] - minPrice
			// Update maximum profit if current profit is higher
			if currentProfit > maxProfit {
				maxProfit = currentProfit
			}
		}
	}

	return maxProfit
}

// Test cases
func main() {
	// Test case 1
	prices1 := []int{7, 1, 5, 3, 6, 4}
	println("Test 1:", maxProfit(prices1)) // Expected: 5

	// Test case 2
	prices2 := []int{7, 6, 4, 3, 1}
	println("Test 2:", maxProfit(prices2)) // Expected: 0
}
