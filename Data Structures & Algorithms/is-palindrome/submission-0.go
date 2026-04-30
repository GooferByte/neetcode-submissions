func isPalindrome(s string) bool {
	i, j := 0, len(s)-1
	for i<j{
		for i<j && !isAlphaNumeric(rune (s[j])){
			j--
		}
		for i<j && !isAlphaNumeric(rune (s[i])){
			i++
		}
		if unicode.ToLower(rune (s[i])) != unicode.ToLower(rune(s[j])){
			return false
		}
		i++
		j--
	}
	return true
}

func isAlphaNumeric(s rune) bool{
	return unicode.IsLetter(s) || unicode.IsDigit(s)
}
