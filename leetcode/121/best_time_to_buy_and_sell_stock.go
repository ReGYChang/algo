package main

func maxProfit(prices []int) int {
	minPrice, profit, l := prices[0], 0, len(prices)
	for i := 0; i < l; i++ {
		if prices[i] < minPrice {
			minPrice = prices[i]
		}

		if diff := prices[i] - minPrice; diff > profit {
			profit = diff
		}
	}

	return profit
}
