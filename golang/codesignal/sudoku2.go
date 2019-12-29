package main

import (
	"fmt"
	//"os"
)

func main() {

	var grid [][]string
	grid = [][]string{
		{".", ".", ".", "1", "4", ".", ".", "2", "."},
		{".", ".", "6", ".", ".", ".", ".", ".", "."},
		{".", ".", ".", ".", ".", ".", ".", ".", "."},

		{".", ".", "1", ".", ".", ".", ".", ".", "."},
		{".", "6", "7", ".", ".", ".", ".", ".", "9"},
		{".", ".", ".", ".", ".", ".", "8", "1", "."},

		{".", "3", ".", ".", ".", ".", ".", ".", "6"},
		{".", ".", ".", ".", ".", "7", ".", ".", "."},
		{".", ".", ".", "5", ".", ".", ".", "7", "."},
	}

	fmt.Println(sudoku2(grid))
}

func sudoku2(grid [][]string) bool {
	var nums []bool
	var nums1 []bool

	for i := 0; i <= 8; i++ {
		nums = make([]bool, 9)
		nums1 = make([]bool, 9)

		for j := 0; j <= 8; j++ {
			if grid[i][j] != "." {
				val := grid[i][j][0] - 49
				if nums[val] {
					fmt.Println("there")
					return false
				} else {
					nums[val] = true
				}

			}

			if grid[j][i] != "." {
				val1 := grid[j][i][0] - 49
				if nums1[val1] {
					fmt.Println("here")
					return false
				} else {
					nums1[val1] = true
				}
			}
		}
	}

	var nums2 []bool

	for pointRow := 0; pointRow < 8; pointRow += 3 {
		for pointCol := 0; pointCol < 8; pointCol += 3 {
			//fmt.Println("point", pointRow, pointCol)
			nums2 = make([]bool, 9)

			row := pointRow
			for i := 0; i < 3; i += 1 {
				col := pointCol
				for j := 0; j < 3; j += 1 {
					//fmt.Println(row, col)

					if grid[row][col] != "." {
						val := grid[row][col][0] - 49
						if nums2[val] {
							fmt.Println("hahaha")
							return false
						} else {
							nums2[val] = true
						}
					}

					col++
				}
				row++
			}
		}
	}

	//fmt.Println(nums)

	return true
}
