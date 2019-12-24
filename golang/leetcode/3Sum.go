package main

import (
	"fmt"
	//"os"
	"sort"
)

func main() {

	nums := []int{}

	nums = []int{1, 1, -2, 0, -1, -5, 2, 7}
	fmt.Println(threeSum(nums))

	/*
		nums = []int{-1, 0, 1, 2, -1, -4}
		fmt.Println(threeSum(nums))

			nums = []int{-1, 0, -1, -2, -1, -4}
			fmt.Println(threeSum(nums))

			nums = []int{0, 0, 0}
			fmt.Println(threeSum(nums))

	*/
}

func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	var res [][]int
	for i := 0; i < len(nums)-2; i++ {

		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		for j, k := i+1, len(nums)-1; j < k; {
			fmt.Println(i, j, k)
			fmt.Println(nums[i], nums[j], nums[k])

			if j > i+1 && nums[j] == nums[j-1] {
				j++
				fmt.Println("h")
				continue
			}

			if k < len(nums)-1 && nums[k] == nums[k+1] {
				k--
				fmt.Println("K")
				continue
			}

			switch sum := nums[i] + nums[j] + nums[k]; {
			case sum == 0:
				v := []int{nums[i], nums[j], nums[k]}
				res = append(res, v)
				j++
				k--
			case sum < 0:
				j++
			default:
				k--
			}
		}
	}
	return res
}
