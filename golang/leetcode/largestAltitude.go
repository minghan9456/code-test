package main

import "fmt"

func main() {
	// gain := []int{-5, 1, 5, 0, -7}
	gain := []int{-4, -3, -2, -1, 4, 3, 2}
	fmt.Println(largestAltitude(gain))
}

func largestAltitude(gain []int) int {
	start := 0
	max := 0

	for i := 0; i < len(gain); i++ {
		start += gain[i]
		if start > max {
			max = start
		}
	}

	return max
}
