package main

import "fmt"

func main() {
	nums := []int{5, 1, 3, 100}
	fmt.Println(rob(nums))
}

func max(a, b int) int {
	if a >= b {
		return a
	} else {
		return b
	}
}

func rob(nums []int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}

	dp2 := 0
	dp1 := 0
	// 5, 5, 8, 105

	for i := 0; i < n; i++ {
		dp := max(dp2+nums[i], dp1)
		dp2 = dp1
		dp1 = dp
	}

	return dp1
}

// [5,1,3,100] => pick 0, 3 => 105
// Two options:
// Take the money if we did't robber house i - 1
// Skip
// TC : O(n)
// SC : O(n)
/*
func rob(nums []int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}

	dp := make([]int, n)
	// 5, 5, 8, 105

	for i := 0; i < n; i++ {
		if i == 0 {
			dp[i] = nums[i]
		} else if i == 1 {
			dp[i] = max(dp[i-1], nums[i])
		} else {
			dp[i] = max(dp[i-2]+nums[i], dp[i-1])
		}
	}

	return dp[len(dp)-1]
}
*/
