package main

import (
	"fmt"
	"strconv"
)

func main() {
	// nums := []int{0, 1, 2, 4, 5, 6}
	// nums := []int{0, 1, 2, 4, 5, 7}
	// nums := []int{0, 2, 3, 4, 6, 8, 9}
	// nums := []int{}
	nums := []int{0}
	fmt.Println(summaryRanges(nums))
}

func summaryRanges(nums []int) []string {
	ans := []string{}
	if len(nums) == 0 {
		return ans
	}

	head := nums[0]
	start := nums[0]
	pos := 1

	for pos < len(nums) {
		if start+1 == nums[pos] {
			start++
		} else {
			if head == start {
				ans = append(ans, strconv.Itoa(head))
			} else {
				ans = append(ans, strconv.Itoa(head)+"->"+strconv.Itoa(start))
			}
			head = nums[pos]
			start = nums[pos]
		}
		pos++
	}

	if head == start {
		ans = append(ans, strconv.Itoa(head))
	} else {
		ans = append(ans, strconv.Itoa(head)+"->"+strconv.Itoa(start))
	}

	return ans
}
