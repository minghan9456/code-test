package main

import (
	"fmt"
)

func main() {
	var A [][]int

	A = [][]int{{1, 1, 0}, {1, 0, 1}, {0, 0, 0}}
	fmt.Println(flipAndInvertImage(A))

	A = [][]int{{1, 1, 0}, {1, 0, 1}, {0, 0, 1}, {1, 1, 0, 1}}
	fmt.Println(flipAndInvertImage(A))

}

func flipAndInvertImage(A [][]int) [][]int {
	for i, nums := range A {
		numsLen := len(nums) - 1
		for j := 0; j <= numsLen/2; j++ {
			nums[j], nums[numsLen-j] = nums[numsLen-j]^1, nums[j]^1
		}

		A[i] = nums
	}

	return A
}
