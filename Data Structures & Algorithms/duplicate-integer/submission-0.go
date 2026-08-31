func hasDuplicate(nums []int) bool {
    // создать мапу, ключ это число из среза, а элемент это кол-ва
	a := make(map[int]int)
	for _, num := range nums {
		a[num]++
		if a[num] > 1 {
			return true
		}	
	}
	return false
}
