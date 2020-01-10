package main

import (
	"fmt"
	//"sort"
)

func main() {

	var nums []int

	nums = []int{4, 1, 2, 1, 2}
	fmt.Println(singleNumber(nums))
}

func singleNumber(nums []int) int {

	ret := nums[0]

	for i := 1; i <= len(nums)-1; i++ {
		ret ^= nums[i]
		fmt.Println(ret)
	}

	return ret
}
