func maxArea(heights []int) int {
	max := 0
	left, right := 0, len(heights)-1
	for left < right {
		lower := heights[right]
		if heights[left] < heights[right] {
			lower = heights[left]
		}
		if max < lower * (right - left) {
			max = lower * (right - left)
		}
		if heights[left] < heights[right] {
			left++
		} else {
			right--
		}
	}
	return max
}
