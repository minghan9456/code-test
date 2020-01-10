package main

import (
	"fmt"
)

func main() {
	var nums []int

	nums = []int{1, 2, 3, 2}
	fmt.Println(containsDuplicate(nums))

	nums = []int{1, 2, 3, 4}
	fmt.Println(containsDuplicate(nums))

	nums = []int{1, 1, 1, 3, 3, 4, 3, 2, 4, 2}
	fmt.Println(containsDuplicate(nums))
}

func containsDuplicate(nums []int) bool {
	if len(nums) <= 1 {
		return false
	}

	j := 0
	for j <= len(nums)-1 {
		tmp := nums[j]

		for i := j + 1; i <= len(nums)-1; i++ {
			if tmp == nums[i] {
				return true
			}
		}

		j++
	}

	return false
}
