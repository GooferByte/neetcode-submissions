func maxArea(heights []int) int {
	maxCap :=0
	left, right :=0, len(heights)-1
	for right>left{
		cap :=0
		length := 0
		breadth := right-left
		if heights[left]<=heights[right]{
			length= heights[left]
			left++
		} else {
			length = heights[right]
			right--
		}
		cap = length*breadth
		if cap>maxCap {
			maxCap=cap
		}
	}
	return maxCap
}
