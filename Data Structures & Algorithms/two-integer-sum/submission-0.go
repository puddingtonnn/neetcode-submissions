func twoSum(nums []int, target int) []int {
	a := make(map[int]int)
    for i, v := range nums {
		secNum := target - v
		if idx, ok := a[secNum]; ok {
			return []int{idx, i}
		}
		a[v] = i
	}
	return nil
}
