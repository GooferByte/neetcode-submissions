func longestConsecutive(nums []int) int {
	if len(nums)==0 {return 0}
	set := make(map[int]bool)
	max :=1
	for _, i := range(nums){
		set[i]=true
	}
	for _, i := range(nums){
		if !set[i-1]{
			continue
		}
		curr := i
		cnt := 1
		for set[curr-1]{
			cnt++
			curr--
		}
		if (cnt>max){
			max = cnt
		} 
	}
	return max
}
