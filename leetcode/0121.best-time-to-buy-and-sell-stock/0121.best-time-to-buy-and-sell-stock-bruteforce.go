package leetcode

// Brute force
func maxProfit(prices []int) int {

	maxProfit := 0
	for i := 0; i < len(prices); i++ {
		for j := i + 1; j < len(prices); j++ {
			if prices[i] < prices[j] && prices[j]-prices[i] > maxProfit {
				maxProfit = prices[j] - prices[i]
			}

		}
	}
	return maxProfit
}
