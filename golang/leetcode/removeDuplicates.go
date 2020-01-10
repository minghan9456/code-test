package main

import (
	"fmt"
)

func main() {

	var nums []int

	//nums = []int{0}
	nums = []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4, 5, 5, 5, 5, 5, 6, 6, 6, 7, 7, 8, 8, 8, 9, 9, 9, 9, 9}
	fmt.Println("h", nums)
	fmt.Println(removeDuplicates(nums))
}

func removeDuplicates(nums []int) int {

	if nums == nil || len(nums) == 0 {
		return len(nums)
	}

	tmp := nums[0]
	i := 1
	count := 0

	for i < len(nums) {
		if tmp != nums[i] {
			tmp = nums[i]
			count++
		} else {
			fmt.Println(i)
			fmt.Println(nums)
			nums = append(nums[:i], nums[i+1:]...)
			fmt.Println(nums)
			i = count
		}

		i++
	}

	return len(nums)
}
