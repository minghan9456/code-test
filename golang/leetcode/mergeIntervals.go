package main

import (
	"fmt"
	"sort"
)

func main() {
	// i := [][]int{{1, 3}, {2, 6}, {8, 10}, {15, 18}}
	i := [][]int{{4, 5}, {1, 4}, {8, 10}, {9, 18}, {9, 20}}
	// i := [][]int{{1, 4}, {4, 5}}
	// i := [][]int{{1, 4}, {0, 4}}
	// i := [][]int{{0, 4}, {1, 4}}
	// i := [][]int{{1, 4}, {1, 5}}
	// i := [][]int{{2, 4}, {1, 4}}
	fmt.Println(merge(i))
}

func merge(intervals [][]int) [][]int {
	// Time: O(nlogn)
	sort.Slice(intervals, func(a, b int) bool {
		return (intervals[a][0] < intervals[b][0]) || ((intervals[a][0] == intervals[b][0]) && (intervals[a][1] < intervals[b][1]))
	})

	// Space: O(n)
	r := [][]int{}
	r = append(r, intervals[0])

	rIdx, idx := 0, 1

	// Time: O(n)
	for idx <= len(intervals)-1 {
		rEnd := r[rIdx][1]
		start2, end2 := intervals[idx][0], intervals[idx][1]

		if rEnd < start2 {
			r = append(r, intervals[idx])
			rIdx++
		}

		if end2 > rEnd {
			r[rIdx][1] = end2
		} else {
			r[rIdx][1] = rEnd
		}

		idx++
	}

	return r
}

/*
func max(x, y int) int {
	if x >= y {
		return x
	} else {
		return y
	}
}

func merge(intervals [][]int) [][]int {
	sort.Slice(intervals, func(a, b int) bool {
		return (intervals[a][0] < intervals[b][0]) || ((intervals[a][0] == intervals[b][0]) && (intervals[a][1] < intervals[b][1]))
	})

	START, END := 0, 1
	r := [][]int{}

	for _, curr := range intervals {
		if len(r) == 0 || r[len(r)-1][END] < curr[START] {
			r = append(r, curr)
		} else {
			r[len(r)-1][END] = max(curr[END], r[len(r)-1][END])
		}
	}

	return r
}
*/
