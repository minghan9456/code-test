package main

import (
	"fmt"
)

func main() {
	target := 7
	nums := []int{2, 3, 1, 2, 4, 3}
	fmt.Println(minSubArrayLen(target, nums))
}

func min(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}

// TC : O(n)
// SC : O(1)
// Two Pointers
func minSubArrayLen(target int, nums []int) int {
	flag := false

	l := 0
	r := 0
	n := len(nums)
	minLen := n
	sum := 0

	for r < n {
		sum += nums[r]

		for sum >= target {
			flag = true
			minLen = min(minLen, r-l+1)

			sum -= nums[l]
			l++
		}

		r++
	}

	if !flag {
		return 0
	} else {
		return minLen
	}
}

/*
func minSubArrayLen(target int, nums []int) int {
	flag := false
	minLen := len(nums)

	for i := 0; i < len(nums); i++ {
		sum := 0
		count := 0
		for j := i; j < len(nums); j++ {
			count++
			sum += nums[j]
			if sum >= target {
				flag = true
				minLen = min(minLen, count)
				break
			}
		}
	}

	if !flag {
		return 0
	}

	return minLen
}
*/
