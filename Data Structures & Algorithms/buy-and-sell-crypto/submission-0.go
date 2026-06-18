func maxProfit(prices []int) int {
	minPrices := 100
	maxProfit := 0
	for i := 0; i < len(prices); i++ {
		if prices[i] < minPrices {
			minPrices = prices[i]
		}
		if i > 0 && prices[i] - minPrices > maxProfit {
			maxProfit = prices[i] - minPrices
		}
	}
	return maxProfit
}
