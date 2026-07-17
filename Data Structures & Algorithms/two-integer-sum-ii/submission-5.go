func twoSum(numbers []int, target int) []int {
	m := make(map[int]int)
	for i, val := range(numbers){
		if _, ok := m[target-val]; ok{
			return []int{m[target-val], i+1}
		}
		m[val]=i+1
	}
	return []int{0,0}
}
