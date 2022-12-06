package main

import "fmt"

func findMaxTradingCost(input []int) int {
	if len(input) == 0 {
		return 0
	}
	maxProfit := 0
	minProfit := input[0]
	for i := 1; i < len(input); i++ {
		if input[i] < minProfit {
			minProfit = input[i]
		} else {
			if input[i]-minProfit > maxProfit {
				maxProfit = input[i] - minProfit
			}
		}	
	}
	return maxProfit
}

func main() {
	stockPrices := []int{100, 180, 260, 310, 40, 535, 695}
	fmt.Println(findMaxTradingCost(stockPrices))
}
