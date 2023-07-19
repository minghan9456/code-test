package main

import "fmt"

func main() {
	fmt.Println("vim-go")
}

/*
As the rows are sorted we can use binary search to find the index of 1st negatve element.

Find the index of 1st negative element then simply (total number of elements in the row - index) would be the negative numbers in that row.
Do this for each row and return the total answer.

TC : O(mlogn)
SC : O(1)
*/

func countNegatives(grid [][]int) int {
	negatives := 0
	m := len(grid)
	n := len(grid[0])

	for i := 0; i < m; i++ {
		pos := bs(grid[i])
		if pos != -1 {
			negatives += n - pos
		}
	}

	return negatives
}

func bs(arr []int) int {
	rtn := -1
	l := 0
	r := len(arr) - 1

	for l <= r {
		mid := (r-l)/2 + l
		if arr[mid] < 0 {
			rtn = mid
			r = mid - 1
		} else {
			l = mid + 1
		}
	}

	return rtn
}

/*
The key idea is to take advantage of the sorted nature of the matrix, where both the rows and columns are sorted in non-increasing order. This means that once we find a negative number in a row, all the following numbers in that row will also be negative. Similarly, if a cell at the top of a column is negative, all the cells below it in that column will also be negative.

Negative numbers will form a staircase. So we start from the bottom-left, and follow along the staircase. For each row, we find the edge of the staircase, add all the count of elements to the right to the answer, and then move up a row and repeat.
TC : O(m+n)
SC : O(1)
*/
/*
func countNegatives(grid [][]int) int {
	negatives := 0
	rows := len(grid) - 1
	cols := len(grid[0]) - 1

	startRow := rows
	startCol := 0

	for startRow >= 0 {
		for startCol <= cols && grid[startRow][startCol] >= 0 {
			startCol++
		}

		negatives += cols - startCol + 1
		startRow--
	}

	return negatives
}
*/
