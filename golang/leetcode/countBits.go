package main

import (
	"fmt"
)

func main() {

	var num int

	num = 12
	fmt.Println(countBits(num))
}

func countBits(num int) []int {
	var ret []int
	ret = make([]int, num+1)

	for i := 1; i <= num; i++ {
		ret[i] = ret[i&(i-1)] + 1
	}

	return ret
}
