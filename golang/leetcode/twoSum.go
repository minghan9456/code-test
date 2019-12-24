package main

import (
	"fmt"
)

func main() {
	var nums []int
	var target int

	/*
		nums = []int{2, 7, 11, 15}
		target = 9
		fmt.Println(twoSum(nums, target))
	*/

	nums = []int{0, 3, 4, 0}
	target = 0
	fmt.Println(twoSum(nums, target))

}

func twoSum(nums []int, target int) []int {
	var clct []int

	for i, num := range nums {
		for j, c_num := range clct {

			if c_num == target-num && i != j {
				if i > j {
					return []int{j, i}
				} else {
					return []int{i, j}
				}
			}
		}

		clct = append(clct, num)
	}

	return []int{}
}

/*
func twoSum(nums []int, target int) []int {
	for i, num := range nums {
		for j, j_num := range nums {
			if j_num == target-num && i != j {
				return []int{i, j}
			}
		}
	}

	return []int{}
}
*/
