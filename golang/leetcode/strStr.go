package main

import "fmt"

func main() {
	// fmt.Println(strStr("sadbutsad", ""))
	// fmt.Println(strStr("s", "sadbutsad"))
	// fmt.Println(strStr("sadbutsad", "sad"))
	// fmt.Println(strStr("butsad", "sad"))
	fmt.Println(strStr("leetcode", "leeto"))
}

func strStr(haystack string, needle string) int {

	hLen := len(haystack)
	nLen := len(needle)

	if nLen == 0 {
		return 0
	}

	if hLen < nLen {
		return -1
	}

	for i := 0; i <= hLen-nLen; i++ {
		match := true
		for j := 0; j <= nLen-1; j++ {
			fmt.Println(string(needle[j]), string(haystack[i+j]))
			if needle[j] != haystack[i+j] {
				match = false
				break
			}
		}

		fmt.Println(match)
		if match {
			return i
		}
	}

	return -1
}
