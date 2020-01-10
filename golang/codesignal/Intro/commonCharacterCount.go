package main

import "fmt"

func main() {
	var s1, s2 string

	s1 = "abca"
	s2 = "xyzbac"

	fmt.Println(commonCharacterCount(s1, s2))
}

func commonCharacterCount(s1 string, s2 string) int {

	clct := make([]bool, len(s2))

	count := 0

	for i := 0; i <= len(s1)-1; i++ {
		for j := 0; j <= len(s2)-1; j++ {
			if !clct[j] && s1[i] == s2[j] {
				clct[j] = true
				count++
				break
			}
		}
	}

	return count
}
