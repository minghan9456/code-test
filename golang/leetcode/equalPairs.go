package main

import "fmt"

func main() {
	grid := [][]int{{3, 2, 1}, {1, 7, 6}, {2, 7, 7}}
	fmt.Println(equalPairs(grid))
}

// hash map
// trie

func equalPairs(grid [][]int) int {
	n := len(grid)

	count := 0

	for row := 0; row < n; row++ {
		for col := 0; col < n; col++ {
			match := true
			fmt.Println("------")
			for i := 0; i < n; i++ {
				match = true
				fmt.Println("c", row, i, i, col, grid[row][i], grid[i][col], grid[row][i] == grid[i][col])
				if grid[row][i] != grid[i][col] {
					match = false
					// break
				}
			}

			if match == true {
				fmt.Println(row, col)
				count++
			}
		}
	}

	return count
}

/*
0, 0 0, 0
0, 1 1, 0
0, 2 2, 0
0, 3 3, 0

0, 0 0, 1
0, 1 1, 1
0, 2 2, 1
0, 3 3, 1

0, 0 0, 2
0, 1 1, 2
0, 2 2, 2
0, 3 3, 2
*/
