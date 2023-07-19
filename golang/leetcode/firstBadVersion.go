package main

import "fmt"

func main() {
	fmt.Println("vim-go")
}

/**
 * Forward declaration of isBadVersion API.
 * @param   version   your guess about first bad version
 * @return 	 	      true if current version is bad
 *			          false if current version is good
 * func isBadVersion(version int) bool;
 */

// 4
// 1 2 3 4 5
// 3

// BINARY SEARCH
// TC : O(logn)
// SC : O(1)
func firstBadVersion(n int) int {
	start, end, mid := 0, n, 0

	for start <= end {
		mid = start + (end-start)/2
		fmt.Println(start, mid, end)

		if isBadVersion(mid) {
			fmt.Println(mid, "bad", isBadVersion(mid-1))
			if !isBadVersion(mid - 1) {
				return mid
			}

			end = mid - 1
		} else {
			start = mid + 1
		}
	}

	return mid
}

// TC : O(n/2) => O(n)
// SC : O(1)
func firstBadVersion(n int) int {
	i := n/2 + 1

	for 0 < i && i <= n {
		if isBadVersion(i) {
			if !isBadVersion(i - 1) {
				return i
			}
			i--
		} else {
			i++
		}
	}

	return 1
}
