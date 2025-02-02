package main

func containsDuplicate(nums []int) bool {
	numRecord := make(map[int]bool)
	for _, num := range nums {
		if _, ok := numRecord[num]; ok {
			return true
		}
		numRecord[num] = true
	}

	return false
}
