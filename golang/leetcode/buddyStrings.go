package main

import "fmt"

func main() {
	s := "aaaaaaabc"
	goal := "aaaaaaacb"
	fmt.Println(buddyStrings(s, goal))
}

/*
If A.length() != B.length(): no possible swap

If A == B, we need swap two same characters. Check is duplicated char in A.

In other cases, we find index for A[i] != B[i]. There should be only 2 diffs and it's our one swap.
*/

func buddyStrings(s string, goal string) bool {
	if len(s) != len(goal) {
		return false
	}

	hset := make(map[string]int)
	if s == goal {
		for i := 0; i < len(s); i++ {
			char := string(s[i])
			hset[char]++

			if hset[char] >= 2 {
				return true
			}
		}
	}

	diffs := []int{}
	for i := 0; i < len(s); i++ {
		if s[i] != goal[i] {
			diffs = append(diffs, i)
		}
	}

	return len(diffs) == 2 && (s[diffs[0]] == goal[diffs[1]] && s[diffs[1]] == goal[diffs[0]])
}
