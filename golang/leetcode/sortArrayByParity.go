package main

import (
	"fmt"
)

func main() {
	var A []int

	A = []int{3, 1, 2, 4}
	fmt.Println(sortArrayByParity(A))

	A = []int{4, 2, 7, 5, 4, 8}
	fmt.Println(sortArrayByParity(A))

	A = []int{4, 2, 7, 5, 4, 8, 1}
	fmt.Println(sortArrayByParity(A))

}

func sortArrayByParity(A []int) []int {
	headIdx := 0
	tailIdx := len(A) - 1

	for headIdx < tailIdx {
		if A[headIdx]%2 == 0 {
			headIdx++
		} else {
			if A[tailIdx]%2 != 0 {
				tailIdx--
			} else {
				if A[tailIdx]%2 == 0 {
					A[headIdx], A[tailIdx] = A[tailIdx], A[headIdx]
					headIdx++
					tailIdx--
				}
			}
		}
	}

	return A
}

/*
func sortArrayByParity(A []int) []int {

	aLen := len(A)
	ret := make([]int, aLen)

	headIdx, tailIdx := 0, aLen-1
	for _, num := range A {
		if num%2 == 0 {
			ret[headIdx] = num
			headIdx++
		} else {
			ret[tailIdx] = num
			tailIdx--
		}
	}

	return ret
}
*/
