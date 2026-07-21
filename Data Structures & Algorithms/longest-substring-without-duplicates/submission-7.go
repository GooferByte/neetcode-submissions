func lengthOfLongestSubstring(s string) int {
	m := make(map[byte]int)
	max :=0
	count :=0
	left :=0
	for right:=0; right<len(s); right++{
		c := s[right]
		if idx, ok := m[c]; ok && left<=idx{
			left=idx+1
		}
		m[c]=right
		count= right-left+1
		if count > max {
			max= count
		}
	}
	return max
}
