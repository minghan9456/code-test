package main

import (
	"fmt"
)

func main() {

	var A []int

	A = []int{4, 2, 5, 7}
	fmt.Println(sortArrayByParityII(A))
}

func sortArrayByParityII(A []int) []int {
	j := 1

	for i := 0; i < len(A); i += 2 {
		if A[i]%2 == 1 {
			for A[j]%2 == 1 {
				j += 2
			}

			A[i], A[j] = A[j], A[i]
		}
	}

	return A
}

/*
func sortArrayByParityII(A []int) []int {

	ret := make([]int, len(A))

	t := 0
	for i := 0; i < len(A); i += 1 {
		if A[i]%2 == 0 {
			ret[t] = A[i]
			t += 2
		}
	}

	t = 1
	for i := 0; i < len(A); i += 1 {
		if A[i]%2 == 1 {
			ret[t] = A[i]
			t += 2
		}
	}

	return ret
}
*/
