package main

import (
	"fmt"
)

func main() {
	var k int
	var nums []int

	k = 3
	nums = []int{1, 2, 3, 4, 5, 6, 7}

	rotate(nums, k)
	fmt.Println(nums)
}

func rotate(nums []int, k int) {

	length := len(nums) - 1

	for k > 0 {
		tmp := nums[length]

		for i := length; i >= 1; i-- {
			nums[i] = nums[i-1]
		}
		nums[0] = tmp

		k--
	}

}
