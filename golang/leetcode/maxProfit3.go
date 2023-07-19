package main

import "fmt"

func main() {
	prices := []int{1, 3, 2, 8, 4, 9}
	fmt.Println(maxProfit(prices, 2))
}

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func maxProfit(prices []int, fee int) int {
	n := len(prices)
	sold := make([]int, n)
	hold := make([]int, n)

	hold[0] = -prices[0]

	for i := 1; i < n; i++ {
		sold[i] = max(sold[i-1], hold[i-1]+prices[i]-fee)
		hold[i] = max(hold[i-1], sold[i-1]-prices[i])
	}

	fmt.Println(sold)
	fmt.Println(hold)

	return sold[n-1]
}
