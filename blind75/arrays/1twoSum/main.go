package main

func twoSum(nums []int, target int) []int {
	numRecord := make(map[int]int)

	for i, num := range nums {
		compliment := target - num
		if _, found := numRecord[compliment]; found {
			return []int{numRecord[compliment], i}
		} else {
			numRecord[num] = i
		}
	}

	return []int{}
}
