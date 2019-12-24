package main

import (
	"fmt"
)

func main() {

	var A []int

	A = []int{-4, -1, 0, 3, 10}
	fmt.Println(sortedSquares(A))
}

func sortedSquares(A []int) []int {
	ret := make([]int, 0)

	l, r := 0, len(A)-1
	for l <= r {
		left, right := abs(A[l]), abs(A[r])
		if left > right {
			ret = append([]int{int(left * left)}, ret...)
			l++
		} else {
			ret = append([]int{int(right * right)}, ret...)
			r--
		}
	}

	return ret
}

func abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}

/*
func sortedSquares(A []int) []int {

	for i := 0; i < len(A); i++ {
		A[i] = A[i] * A[i]
	}

	l := 0
	r := len(A) - 1
	for l < r {
		for i := l; i < r; i++ {
			if A[i] > A[i+1] {
				A[i], A[i+1] = A[i+1], A[i]
			}
		}
		r--

		for i := r; i > l; i-- {
			if A[i-1] > A[i] {
				A[i-1], A[i] = A[i], A[i-1]
			}
		}
		l++
	}

	return A
}
*/
