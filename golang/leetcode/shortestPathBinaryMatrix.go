package main

import "fmt"

func main() {
	grid := [][]int{[]int{0, 0, 0}, []int{1, 1, 0}, []int{1, 1, 0}}
	// grid := [][]int{[]int{1, 0, 0}, []int{1, 1, 0}, []int{1, 1, 0}}
	fmt.Println(shortestPathBinaryMatrix(grid))
}

func shortestPathBinaryMatrix(grid [][]int) int {
	n := len(grid)
	queue := [][3]int{{1, 0, 0}}                                                           // step, row, column
	dirs := [][2]int{{1, 0}, {1, 1}, {0, 1}, {-1, -1}, {-1, 0}, {0, -1}, {1, -1}, {-1, 1}} // row, column

	// top-left [0][0]int is not 0 => return -1
	// bottom-right [n-1][n-1]int is not 0 => return -1
	if grid[0][0] == 1 || grid[n-1][n-1] == 1 {
		return -1
	}

	// for init first step
	grid[0][0] = 1

	for len(queue) > 0 {
		step, row, col := queue[0][0], queue[0][1], queue[0][2]
		queue = queue[1:]

		// To End
		if row == n-1 && col == n-1 {
			return step
		}

		// find 8 dir can keep go
		for _, move := range dirs {
			nRow, nCol := row+move[0], col+move[1]

			if (0 <= nRow && nRow < n) && (0 <= nCol && nCol < n) && (grid[nRow][nCol] == 0) {
				grid[nRow][nCol] = 1
				queue = append(queue, [3]int{step + 1, nRow, nCol})
			}
		}
	}

	return -1
}
