package main

import (
	"fmt"
)

func main() {

	var A []int

	A = []int{1, 2, 3, 3}
	fmt.Println(repeatedNTimes(A))

	A = []int{2, 1, 2, 5, 3, 2}
	fmt.Println(repeatedNTimes(A))

	A = []int{5, 1, 5, 2, 5, 3, 5, 4}
	fmt.Println(repeatedNTimes(A))
}

func repeatedNTimes(A []int) int {
	length := len(A)
	clct := make([]int, 10000)

	if length%2 != 0 {
		return 0
	}

	for i := 0; i <= length-1; i++ {
		clct[A[i]]++
		if clct[A[i]] == length/2 {
			return A[i]
		}
	}

	return 0
}
