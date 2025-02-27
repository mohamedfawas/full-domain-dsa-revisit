package main

func minEatingSpeed(piles []int, h int) int {
	maxPile := 0

	for _, pile := range piles {
		if pile > maxPile {
			maxPile = pile
		}
	}

	left, right := 1, maxPile

	for left < right {
		mid := left + (right-left)/2

		if canEatAll(piles, h, mid) {
			right = mid
		} else {
			left = mid + 1
		}
	}

	return left

}

func canEatAll(piles []int, h, k int) bool {
	hoursNeeded := 0

	for _, pile := range piles {
		hoursNeeded += (pile + k - 1) / k
	}

	return hoursNeeded <= h
}
