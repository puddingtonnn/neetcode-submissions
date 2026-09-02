func topKFrequent(nums []int, k int) []int {
	count := make(map[int]int)

	for _, num := range nums {
		count[num]++
	}

	buckets := make([][]int, len(nums)+1)

	for num, freq := range count {
		buckets[freq] = append(buckets[freq], num)
	}

	result := make([]int, 0, k)

	for freq := len(buckets)-1; freq >=1; freq-- {
		for _, num := range buckets[freq] {
			result = append(result, num)

			if len(result) == k {
				return result
			}
		}
	}

	return result
}
