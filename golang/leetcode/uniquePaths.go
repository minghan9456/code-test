// https://leetcode.com/problems/unique-paths/

func uniquePaths(m int, n int) int {
    if m == 0 || n == 0 {
        return 0
    }
    
    if m == 1 || n == 1 {
        return 1
    }
    
    cur := make([]int, n)
    for i := range cur {
        cur[i] =  1
    }
    
    for i := 1; i < m; i++ {
        for j := 1; j < n; j++ {
            cur[j] += cur[j-1]
        }
        fmt.Println(i, cur)
    }
    
    // fmt.Println(cur)
    
    /*
        1 [1 2 3 4 5 6 7]
        2 [1 3 6 10 15 21 28]
          [1 3 6 10 15 21 28]
    */
    
    return cur[len(cur) - 1]
    
    /*
    dp := make([][]int, m)
    
    for i := range dp {
        dp[i] = make([]int, n)
        dp[i][0] = 1
    }
    
    for i := range dp[0] {
        dp[0][i] = 1
    }
    
    for i := 1; i < m; i ++{
        for j := 1; j < n; j++ {
            dp[i][j] = dp[i-1][j] + dp[i][j-1]    
        }
    }
    
    // fmt.Println(dp)
    
    // [
    // [1 1 1 1 1 1 1]
    // [1 2 3 4 5 6 7] 
    // [1 3 6 10 15 21 28]
    // ]
    
    return dp[m -1][n - 1]
    */
}
