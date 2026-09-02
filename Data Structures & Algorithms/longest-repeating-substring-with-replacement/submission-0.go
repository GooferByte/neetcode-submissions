func characterReplacement(s string, k int) int {
	if k>= len(s){
		return len(s)
	}
	count := [26]int{}
	left := 0
	ans := 0
	maxFreq := 0
	for right:=0; right<len(s); right++{
		count[s[right]- 'A']++
		maxFreq = max(count[s[right]- 'A'], maxFreq) 
		if (right-left+1)- maxFreq >k {
			count[s[left]-'A']--
			left++
		}
		ans = max(ans, right-left+1)
	}
	return ans
}

func max(i, j int) int {
	if i>j {
		return i
	}
	return j
}
