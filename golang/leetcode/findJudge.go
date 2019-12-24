package main

import (
	"fmt"
)

func main() {

	var N int
	var trust [][]int

	N = 4
	trust = [][]int{{1, 3}, {1, 4}, {2, 3}, {2, 4}, {4, 3}}
	fmt.Println(findJudge(N, trust))

	N = 1
	trust = [][]int{}
	fmt.Println(findJudge(N, trust))
}

func findJudge(N int, trust [][]int) int {
	clct := make([][]int, N+1)
	theJudge := make([]int, 0)

	if len(trust) == 0 {
		return N
	}

	for _, t := range trust {
		clct[t[1]] = append(clct[t[1]], t[0])
	}

	for i := 1; i <= N; i++ {
		trusted := false
		for j := 1; j <= N; j++ {
			if i != j {
				if inSlice(j, clct[i]) {
					trusted = true
				} else {
					trusted = false
					break
				}
			}
		}

		if trusted {
			theJudge = append(theJudge, i)
		}
	}

	if len(theJudge) != 1 {
		return -1
	}

	isJudge := true
	for i := 1; i <= N; i++ {
		if inSlice(theJudge[0], clct[i]) {
			isJudge = false
			break
		}
	}

	if isJudge {
		return theJudge[0]
	} else {
		return -1
	}
}

func inSlice(a int, list []int) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}
