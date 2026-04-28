func hasDuplicate(nums []int) bool {
    var freq = make(map[int]int)
	for _, i := range(nums){
		freq[i]++
		if freq[i]>1 {
			return true
		}
	}
	return false
}
