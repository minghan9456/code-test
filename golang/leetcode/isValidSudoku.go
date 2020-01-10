package main

import (
	"fmt"
)

func main() {

	var board [][]byte

	/*
		board = [][]byte{
			{'8', '3', '.', '.', '7', '.', '.', '.', '.'},
			{'6', '.', '.', '1', '9', '5', '.', '.', '.'},
			{'.', '9', '8', '.', '.', '.', '.', '6', '.'},
			{'1', '.', '.', '.', '6', '.', '.', '.', '3'},
			{'4', '.', '.', '8', '.', '3', '.', '.', '1'},
			{'7', '.', '.', '.', '2', '.', '.', '.', '6'},
			{'.', '6', '.', '.', '.', '.', '2', '8', '.'},
			{'.', '.', '.', '4', '1', '9', '.', '.', '5'},
			{'.', '.', '.', '.', '8', '.', '.', '7', '9'},
		}
	*/

	board = [][]byte{
		{'.', '.', '.', '.', '.', '.', '.', '.', '.'},
		{'.', '.', '.', '3', '.', '.', '5', '.', '.'},
		{'.', '.', '.', '.', '.', '.', '.', '.', '.'},
		{'.', '.', '.', '8', '.', '.', '.', '.', '.'},
		{'.', '.', '.', '.', '1', '1', '6', '.', '.'},
		{'.', '.', '.', '.', '.', '.', '.', '.', '.'},
		{'.', '.', '.', '.', '.', '.', '1', '.', '.'},
		{'.', '.', '.', '.', '.', '.', '.', '.', '7'},
		{'.', '.', '.', '.', '.', '.', '.', '4', '.'},
	}

	fmt.Println(isValidSudoku(board))
}

func isValidSudoku(board [][]byte) bool {
	// 0 => 48

	var i int = 0

	i = 0
	for i <= 8 {
		j := 0
		clct1 := make([]bool, 9)
		clct2 := make([]bool, 9)
		for j <= 8 {
			if board[i][j] >= 49 {
				val1 := board[i][j] - 49
				if clct1[val1] == false {
					clct1[val1] = true
				} else {
					//fmt.Println("r", i, j)
					return false
				}
			}

			if board[j][i] >= 49 {
				val2 := board[j][i] - 49
				if clct2[val2] == false {
					clct2[val2] = true
				} else {
					//fmt.Println("c", i, j)
					return false
				}
			}
			j += 1
		}
		i += 1
	}

	i = 0
	for i <= 8 {
		j := 0
		for j <= 8 {
			//fmt.Println("hrer", i, j)
			row := i
			col := j

			clct := make([]bool, 9)

			row_count := 0
			for row_count <= 2 {
				col_count := 0
				for col_count <= 2 {
					//fmt.Println(row+row_count, col+col_count)
					tmpR := row + row_count
					tmpC := col + col_count

					if board[tmpR][tmpC] >= 49 {
						val := board[tmpR][tmpC] - 49
						if clct[val] == false {
							clct[val] = true
						} else {
							//fmt.Println("3x3", tmpR, tmpC)
							return false
						}
					}

					col_count++
				}
				row_count++
			}

			j += 3
		}
		i += 3
	}

	return true
}
