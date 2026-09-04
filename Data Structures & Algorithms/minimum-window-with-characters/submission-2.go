func minWindow(s string, t string) string {
    if t==""{
		return ""
	}
	freqT := make(map[rune]int)
	for _, c := range(t){
		freqT[c]++
	}
	have, need := 0, len(freqT)
	res := []int{-1, -1}
	resLen := math.MaxInt32
	l := 0
	window := make(map[rune]int)

	for r := 0; r<len(s); r++{
		c:= rune(s[r])
		window[c]++

		if freqT[c]>0 && window[c] == freqT[c] {
			have++
		}

		for have == need {
			if (r -l +1) < resLen{
				res = []int{l, r}
				resLen= r-l+1
			}
			window[rune(s[l])]--
			if freqT[rune(s[l])] > 0 && window[rune(s[l])] < freqT[rune(s[l])] {
				have--
			}
			l++
		}
	}
	if res[0] == -1 {
		return ""
	}
	return s[res[0]:res[1]+1]
}
