// https://leetcode.com/problems/best-time-to-buy-and-sell-stock-ii/

/*
// O(n)
func maxProfit(prices []int) int {
    profit := 0
    
    for i := 0; i < len(prices) - 1; i++ {
        if prices[i+1] > prices[i] {
            profit = profit +  prices[i+1] - prices[i]
        }
    }
    
    return profit
}
*/

// DP
func maxProfit(prices []int) int {
    l := len(prices)
    dp := make(map[int]int);
    dp[0] = 0

    for i := 1; i < l; i++ {
        if dp[i - 1] > dp[i - 1] + (prices[i] - prices[i - 1]) {
            dp[i] = dp[i - 1]
        } else {
            dp[i] = dp[i - 1] + (prices[i] - prices[i - 1])
        }
        // fmt.Println(dp[i])
    }

    return dp[l - 1];
};
