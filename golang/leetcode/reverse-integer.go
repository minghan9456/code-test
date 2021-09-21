// https://leetcode.com/problems/reverse-integer/

import "math"

func reverse(x int) int {
    result := 0
    
    if x == 0 {
        return x
    }
    
    for x != 0 {
        tail := x % 10
        newR := result * 10 + tail
        if ((newR - tail) / 10 != result || newR > math.MaxInt32 || newR < math.MinInt32) { 
            return 0
        }
        result = newR
        x /= 10
    }
    
    return result
}
