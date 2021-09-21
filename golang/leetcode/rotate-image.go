// https://leetcode.com/problems/rotate-image/solution/

func rotate(matrix [][]int)  {
    l := len(matrix) - 1
    m := l/2
    if l >= 1 {
        for i := 0; i <= m; i++ {
            matrix[i], matrix[l-i] = matrix[l-i], matrix[i]
        }
        // fmt.Println(matrix)
    
        start := 1
        for i := 0; i <= l; i++ {
            for j := start; j <= l; j++ {
                matrix[i][j], matrix[j][i] = matrix[j][i], matrix[i][j]
            }
            start++
        }
    }
}
