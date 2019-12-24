package main

import (
	"fmt"
)

func main() {

	var arr []int

	arr = []int{1, 2, 2, 1, 1, 3}
	fmt.Println(uniqueOccurrences(arr))

	arr = []int{1, 2}
	fmt.Println(uniqueOccurrences(arr))

	arr = []int{-3, 0, 1, -3, 1, 1, 1, -3, 10, 0}
	fmt.Println(uniqueOccurrences(arr))
}

func uniqueOccurrences(arr []int) bool {
	length := len(arr)
	clct := make([]int, 2000)

	for i := 0; i <= length-1; i++ {
		clct[arr[i]+1000]++
	}

	tmp := make([]bool, 1000)
	for i := 0; i <= len(clct)-1; i++ {
		if clct[i] > 0 {
			if !tmp[clct[i]] {
				tmp[clct[i]] = true
			} else {
				return false
			}
		}
	}

	return true
}
