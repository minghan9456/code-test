package main

import (
	"fmt"
	//"os"
)

func main() {

	var a [][]int
	a = [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}

	fmt.Println(a)
	fmt.Println(rotateImage(a))
}

func rotateImage(a [][]int) [][]int {
	length := len(a) - 1
	mid := length / 2

	for i := 0; i <= mid; i++ {
		a[i], a[length-i] = a[length-i], a[i]
	}

	for x := 0; x <= length; x++ {
		for y := x + 1; y <= length; y++ {
			a[x][y], a[y][x] = a[y][x], a[x][y]
		}
	}

	return a
}
