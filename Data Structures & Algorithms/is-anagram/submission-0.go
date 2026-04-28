func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	var freq = make(map[rune]int)

	for _, v := range(s){
		freq[v]++
	}
	for _, v := range(t){
		freq[v]--
		if freq[v]<0{
			return false
		}
	}
	return true
}

