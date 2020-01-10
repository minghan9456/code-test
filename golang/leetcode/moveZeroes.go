package main

import (
	"fmt"
)

func main() {

	var nums []int

	nums = []int{0, 1, 0, 0, 10, 1, 0, 1, 0, 3, 12}

	fmt.Println(nums)
	moveZeroes(nums)
	fmt.Println(nums)
}

func moveZeroes(nums []int) {
	var j int = 0
	for i := 0; i < len(nums); i++ {
		if nums[i] != 0 {
			fmt.Println(i, j, nums[i], nums[j])
			nums[j], nums[i] = nums[i], nums[j]
			fmt.Println("here", nums)
			j++
		}
	}
}

/*
func moveZeroes(nums []int) {
	if len(nums) == 1 {
		return
	}

	//zero
	p1 := 0
	//none zero
	p2 := len(nums) - 1

	for p2 >= 0 {
		if nums[p1] == 0 {
			fmt.Println(p1, p2)
			nums[p1], nums[p1+1] = nums[p1+1], nums[p1]
		}

		p1++

		if p1 >= p2 {
			p1 = 0
			p2--
		}
	}
}
*/

/*
func moveZeroes(nums []int) {
	start := 0
	end := len(nums) - 1

	for start <= end {
		if nums[start] == 0 {
			fmt.Println(start, end)
			nums[start], nums[end] = nums[end], nums[start]
			end--
		}
		start++
	}
}
*/
