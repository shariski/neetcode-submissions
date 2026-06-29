func minEatingSpeed(piles []int, h int) int {
	maxPiles := 0
	totalPiles := 0
	for _, p := range piles {
		if p > maxPiles {
			maxPiles = p
		}
		totalPiles += p
	}

	left, right := 1, maxPiles
	for left < right {
		mid := left + (right-left)/2
		canFinish := false
		hours := 0
		for _, p := range piles {
			hours += (p + mid - 1) / mid
		}
		canFinish = hours <= h
		if canFinish {
			right = mid
		} else {
			left = mid + 1
		}
	}
	return left
}