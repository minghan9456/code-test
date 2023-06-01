package main

import (
	"fmt"
)

func main() {

	var prices []int

	prices = []int{7, 1, 5, 3, 6, 4}
	fmt.Println(maxProfit(prices))

	prices = []int{7, 6, 4, 3, 1}
	fmt.Println(maxProfit(prices))

	prices = []int{2, 4, 1}
	fmt.Println(maxProfit(prices))
}

func maxProfit(prices []int) int {
	left := 0
	right := 1
	maxP := 0

	for right <= len(prices)-1 {
		if prices[left] < prices[right] {
			if prices[right]-prices[left] > maxP {
				maxP = prices[right] - prices[left]
			}
		} else {
			left = right
		}

		right++
	}

	return maxP
}

/*
func maxProfit(prices []int) int {
	n := len(prices)
	if n < 1 {
		return 0
	}
	dpMin := make([]int, n)
	dpMin[0] = prices[0]
	for i := 1; i < n; i++ {
		if dpMin[i-1] < prices[i] {
			dpMin[i] = dpMin[i-1]
		} else {
			dpMin[i] = prices[i]
		}
	}

	var profit int = 0
	for i := 1; i < n; i++ {
		if prices[i]-dpMin[i-1] > profit {
			profit = prices[i] - dpMin[i-1]
		}
	}

	return profit
}
*/

/*
func maxProfit(prices []int) int {
	var profit int = 0

	for i := 0; i <= len(prices)-1; i++ {
		for j := i + 1; j <= len(prices)-1; j++ {
			minus := prices[j] - prices[i]
			if minus > profit {
				profit = minus
			}
		}
	}

	return profit
}
*/

/*
func maxProfit(prices []int) int {
	var profit, lowPDay int = 0, 0

	for i := 1; i <= len(prices)-1; i++ {
		if prices[lowPDay] > prices[i] {
			lowPDay = i
		}
	}

	minus := 0
	for i := lowPDay + 1; i <= len(prices)-1; i++ {
		minus = prices[i] - prices[lowPDay]
		if minus > profit {
			profit = minus
		}
	}

	return profit
}
*/
