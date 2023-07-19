package main

import (
	"fmt"
	"sort"
)

func main() {
	// {1, 2}, {3, 4}
	// i := [][]int{{1, 2}, {2, 3}, {3, 4}, {1, 3}}
	// [[1 2] [1 3] [2 3] [3 4]]
	i := [][]int{{-52, 31}, {-73, -26}, {82, 97}, {-65, -11}, {-62, -49}, {95, 99}, {58, 95}, {-31, 49}, {66, 98}, {-63, 2}, {30, 47}, {-40, -26}}

	fmt.Println(eraseOverlapIntervals(i))
}

func eraseOverlapIntervals(intervals [][]int) int {
	// Time: O(nlogn)
	// only sort end item
	sort.Slice(intervals, func(a, b int) bool {
		return intervals[a][1] < intervals[b][1]
	})

	end, count := intervals[0][1], 0

	// Time: O(n)
	for i := 1; i < len(intervals); i++ {
		s, e := intervals[i][0], intervals[i][1]

		if s >= end {
			end = e
		} else {
			count++
		}
	}

	return count
}
