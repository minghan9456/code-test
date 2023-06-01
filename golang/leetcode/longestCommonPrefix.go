package main

import (
	"fmt"
)

func main() {
	// strs := []string{"flower", "flow", "flight", "flox"}
	// strs := []string{""}
	strs := []string{"abab", "aba", "abc"}
	// strs := []string{"flower", ""}
	// strs := []string{"dog", "racecar", "car"}
	// strs := []string{"abc", "abcdef", "abc"}
	fmt.Println(longestCommonPrefix(strs))
}

func longestCommonPrefix(strs []string) string {
	prefix := ""
	sLen := len(strs)

	if sLen == 0 {
		return prefix
	} else if sLen == 1 {
		return strs[0]
	}

	x := 0
Exit:
	for y := 0; y <= len(strs[x])-1; y++ {
		isMatch := true
		matchChar := strs[x][y]
		for i := x + 1; i <= len(strs)-1; i++ {
			if y > len(strs[i])-1 {
				break Exit
			}
			if matchChar != strs[i][y] {
				isMatch = false
				break
			}
		}

		if isMatch {
			prefix += string(matchChar)
		} else {
			break
		}
	}

	return prefix
}
