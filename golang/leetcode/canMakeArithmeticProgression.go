package main

import (
	"fmt"
	"sort"
)

func main() {
	// arr := []int{1, 5, 9, 13}
	arr := []int{1, 4, 7, 10}
	11 / 3 = diff
	// arr := []int{1, 4, 7, 11}
	// arr := []int{1, 10, 10, 10, 19}
	// arr := []int{1, 2, 3}
	// arr := []int{1, 3, 5, 7, 9}
	// arr := []int{1, 5, 3}
	fmt.Println(canMakeArithmeticProgression(arr))
}

func canMakeArithmeticProgression(arr []int) bool {
	// O(n log n)
	sort.Ints(arr)

	/*
		arr[x] + arr[n-x] = arr[0] + arr[n]
		TC : O(n/2) => O(n)
	*/

	n := len(arr)
	diff := arr[1] - arr[0]

	for i := 2; i < n; i++ {
		if arr[i]-arr[i-1] != diff {
			return false
		}
	}

	return true
}
