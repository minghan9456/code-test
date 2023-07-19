package main

import (
	"fmt"
)

func main() {
	bombs := [][]int{{1, 2, 3}, {2, 3, 1}, {3, 4, 2}, {4, 5, 3}, {5, 6, 4}}
	fmt.Println(maximumDetonation(bombs))
}

func check(b1, b2 []int) bool {
	dx := b1[0] - b2[0]
	dy := b1[1] - b2[1]
	r := b1[2]
	return (dx*dx+dy*dy <= r*r)
}

// Enumerate the bomb to detonate, and simulate the process using BFS
// Time: O(n^3)
// Space: O(n)
func maximumDetonation(bombs [][]int) int {
	max := 0
	n := len(bombs)

	for i := 0; i < n; i++ {
		// Detone from i, totally detone how much => BFS from i
		count := 0
		queue := []int{i}
		seen := make([]bool, n)
		seen[i] = true

		for len(queue) > 0 {
			count++
			curr := queue[0]
			queue = queue[1:]

			for j := 0; j < n; j++ {
				if check(bombs[curr], bombs[j]) && !seen[j] {
					seen[j] = true
					queue = append(queue, j)
				}
			}
		}

		if count > max {
			max = count
		}
	}

	return max
}
