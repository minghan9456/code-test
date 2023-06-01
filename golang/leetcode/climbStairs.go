package main

import "fmt"

func main() {
	fmt.Println(climbStairs(1))
	fmt.Println(climbStairs(2))
	fmt.Println(climbStairs(3))
	fmt.Println(climbStairs(4))
}

func climbStairs(n int) int {
	dp := make([]int, n+1)
	dp[0] = 1
	dp[1] = 2

	for i := 2; i <= n-1; i++ {
		dp[i] = dp[i-1] + dp[i-2]
	}

	fmt.Println(dp)

	return dp[n-1]
}
