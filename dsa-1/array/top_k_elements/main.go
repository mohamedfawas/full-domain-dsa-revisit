package main

import "sort"

func topKFrequent(nums []int, k int) []int {
	// Count frequency of each number
	freqMap := make(map[int]int)
	for _, num := range nums {
		freqMap[num]++
	}

	// Create a slice of unique numbers
	uniqueNums := make([]int, 0, len(freqMap))
	for num := range freqMap {
		uniqueNums = append(uniqueNums, num)
	}

	// Custom sort based on frequency in descending order
	sort.Slice(uniqueNums, func(i, j int) bool {
		return freqMap[uniqueNums[i]] > freqMap[uniqueNums[j]]
	})

	// Return first k elements
	return uniqueNums[:k]
}
