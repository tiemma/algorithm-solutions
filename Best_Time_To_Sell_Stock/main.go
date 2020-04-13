package maxProfit

// https://leetcode.com/problems/best-time-to-buy-and-sell-stock-ii/

func maxProfit(prices []int) int {
	profit := 0
	for idx, _ := range prices {
		if idx+1 < len(prices) && prices[idx+1] > prices[idx] {
			profit += prices[idx+1] - prices[idx]
		}
	}

	return profit
}
