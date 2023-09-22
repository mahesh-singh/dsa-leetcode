// Sliding window
func maxProfit(prices []int) int {

	maxProfit := 0
	l := 0
	r := 1

	for r < len(prices) {
		if prices[l] > prices[r] {
			//move left, found new low
			l = r
		} else if prices[r]-prices[l] > maxProfit {
			maxProfit = prices[r] - prices[l]
		}
		r = r + 1
	}
	return maxProfit
}