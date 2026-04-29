func groupAnagrams(strs []string) [][]string {
	hash := make(map[string][]string)
	for _, v := range(strs){
		key := sortStr(v)
		hash[key]=append(hash[key], v)
	}
	var res [][]string
	for _, w := range hash{
		res = append(res, w)
	}
	return res
}
func sortStr(str string) string{
	runes := []rune(str)
	sort.Slice(runes, func(i, j int) bool {
		return runes[i]<runes[j]
	})
	return string(runes)
}
