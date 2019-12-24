package main

import (
	"fmt"
)

func main() {

	var n int
	var m int
	var indices [][]int

	n = 2
	m = 3
	indices = [][]int{{0, 1}, {1, 1}}
	fmt.Println(oddCells(n, m, indices))

	n = 2
	m = 2
	indices = [][]int{{1, 1}, {0, 0}}
	fmt.Println(oddCells(n, m, indices))
}

func oddCells(n int, m int, indices [][]int) int {
	matrix := make([][]int, n)

	for i := 0; i < len(matrix); i++ {
		matrix[i] = make([]int, m)
	}

	var rowIdx, colIdx int
	for _, index := range indices {
		rowIdx = index[0]
		colIdx = index[1]

		for ri := 0; ri < len(matrix); ri++ {
			for ci := 0; ci < len(matrix[ri]); ci++ {
				if ri == rowIdx {
					matrix[ri][ci] += 1
				}

				if ci == colIdx {
					matrix[ri][ci] += 1
				}
			}
		}
	}

	var count int = 0
	for ri := 0; ri < len(matrix); ri++ {
		for ci := 0; ci < len(matrix[ri]); ci++ {
			if matrix[ri][ci]%2 != 0 {
				count++
			}
		}
	}

	return count
}
