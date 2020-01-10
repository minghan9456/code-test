package main

import (
	"fmt"
)

func main() {

	var prices []int

	/*
		prices = []int{7, 1, 5, 3, 6, 4}
		fmt.Println(maxProfit(prices))

		prices = []int{1, 2, 3, 4, 5}
		fmt.Println(maxProfit(prices))
	*/

	prices = []int{3, 3}
	fmt.Println(maxProfit(prices))

	//prices = []int{7, 6, 4, 3, 1}
	//fmt.Println(maxProfit(prices))

	//prices = []int{2, 4, 1}
	//fmt.Println(maxProfit(prices))
}

func maxProfit(prices []int) int {
	profit := 0

	peak := 0
	valley := 0

	i := 0
	for i < len(prices)-1 {
		for i < len(prices)-1 && prices[i] >= prices[i+1] {
			i++
		}
		valley = prices[i]
		fmt.Println("valley", valley, i)

		for i < len(prices)-1 && prices[i] <= prices[i+1] {
			i++
		}
		peak = prices[i]
		fmt.Println("peak", peak, i)

		profit += peak - valley
	}

	return profit
}
