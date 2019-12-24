package main

import (
	"fmt"
)

func main() {

	var A []int

	A = []int{0, 1, 0}
	fmt.Println(peakIndexInMountainArray(A))

	A = []int{0, 2, 1, 0}
	fmt.Println(peakIndexInMountainArray(A))

	A = []int{3, 4, 5, 1}
	fmt.Println(peakIndexInMountainArray(A))

	A = []int{3, 4, 4, 6, 10, 5, 1}
	fmt.Println(peakIndexInMountainArray(A))

	A = []int{3, 11, 4, 6, 6, 1, 1}
	fmt.Println(peakIndexInMountainArray(A))
}

func peakIndexInMountainArray(A []int) int {
	for i := 1; i < len(A); i += 1 {
		if A[i-1] < A[i] && A[i] > A[i+1] {
			return i
		}
	}

	return -1
}
