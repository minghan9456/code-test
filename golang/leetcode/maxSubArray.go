package main

import "fmt"

func main() {
	// nums := []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}
	nums := []int{-2, 200, -3, 4, -1, 2, 1, -5, 4}
	fmt.Println(maxSubArray(nums))
}

func maxSubArray(nums []int) int {
	ans := nums[0]
	sum := nums[0]

	for i := 1; i <= len(nums)-1; i++ {
		if sum < 0 {
			sum = nums[i]
		} else {
			sum = sum + nums[i]
		}

		if sum > ans {
			ans = sum
		}
	}

	return ans
}

/*
func maxSubArray(nums []int) int {
	dp := make([]int, len(nums))
	dp[0] = -2

	// f[i] = f[i-1] > 0 ? nums[i] + f[i-1] : nums[i];
	for i := 1; i <= len(nums)-1; i++ {
		if dp[i-1] < 0 {
			dp[i] = nums[i]
		} else {
			dp[i] = dp[i-1] + nums[i]
		}
	}

	max := dp[0]
	for i := 1; i <= len(dp)-1; i++ {
		if dp[i] > max {
			max = dp[i]
		}
	}

	return max
}
*/
