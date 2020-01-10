package main

import (
	"fmt"
)

func main() {
	var a []int

	a = []int{-1, 150, 190, 170, -1, -1, 160, 180}
	fmt.Println("ans", sortByHeight(a))
}

func sortByHeight(a []int) []int {

	j := len(a) - 1
	i := 0
	for j >= 0 {
		i = 0
		for i < j {

			if a[i] >= 0 && a[j] >= 0 && a[i] > a[j] {
				a[i], a[j] = a[j], a[i]
			}

			i++
		}

		j--
	}

	return a
}
