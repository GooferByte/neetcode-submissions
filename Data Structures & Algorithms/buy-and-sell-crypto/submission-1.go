func maxProfit(prices []int) int {
	maxProfit := 0
	buy := prices[0]
	for i:=1; i<len(prices); i++{
		profit := prices[i]-buy
		maxProfit = max(maxProfit, profit)
		buy = min(prices[i], buy)
	}
	return maxProfit
}

func min (a, b int) int {
	if a>b{
		return b
	}
	return a
}

func max(a, b int) int{
	if a>b{
		return a
	}
	return b
}
