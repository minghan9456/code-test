// https://leetcode.com/problems/fibonacci-number/

// reduce space
func fib(n int) int {
    curr := 0
    pTwo := 0
    pOne := 1
    
    if n == 0 {
        return 0
    }
    
    if n == 1 {
        return 1
    }
    
    for i := 1; i < n; i++ {
        curr = pTwo + pOne
        pTwo = pOne
        pOne = curr
    }
    
    return curr
}

/*
BP

func fib(n int) int {
    cache := make(map[int]int)
    cache[0] = 0
    cache[1] = 1
    
    return handler(&cache, n)
}

func handler(cache *map[int]int, num int) int {
    if v, ok := (*cache)[num]; ok {
        return v
    }
    
    return handler(cache, num - 2) + handler(cache, num - 1)
}
*/
