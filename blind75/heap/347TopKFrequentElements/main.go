package main

import "sort"

func topKFrequent(nums []int, k int) []int {
	freqMap := make(map[int]int)

	for _, num := range nums {
		freqMap[num]++
	}

	uniqueNums := []int{}
	for num := range freqMap {
		uniqueNums = append(uniqueNums, num)
	}

	sort.Slice(uniqueNums, func(i, j int) bool {
		return freqMap[uniqueNums[i]] > freqMap[uniqueNums[j]]
	})

	return uniqueNums[:k]
}
