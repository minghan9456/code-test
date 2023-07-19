package main

import "fmt"

func main() {
	fmt.Println("vim-go")
}

/*
binary search

TC : O(mlogn)
SC : O(1)
*/
func nextGreatestLetter(letters []byte, target byte) byte {
	rtn := letters[0]
	l := 0
	r := len(letters) - 1

	for l <= r {
		mid := (r-l)/2 + l
		if letters[mid] > target {
			rtn = letters[mid]
			r = mid - 1
		} else {
			l = mid + 1
		}
	}

	return rtn
}
