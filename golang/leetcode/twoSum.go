package main

import (
	"fmt"
)

func main() {
	var nums []int
	var target int

	nums = []int{2, 7, 11, 15}
	target = 9
	fmt.Println(twoSum(nums, target))

	nums = []int{2, 1, 6, 4, 0}
	target = 5
	fmt.Println(twoSum(nums, target))

	nums = []int{1, 0, 3, 4, 0}
	target = 0
	fmt.Println(twoSum(nums, target))

}

// TC => O(n)
func twoSum(nums []int, target int) []int {
    m := make(map[int]int)
    
    for i, n := range nums{
        j, ok := m[-n]
        m[n - target] = i
        if ok {
            return []int{ j, i }
        }
    }

    return []int{}
}

/*
func twoSum(nums []int, target int) []int {
	var i int = 0
	for i <= len(nums)-1 {
		j := i + 1
		for j <= len(nums)-1 {
			if nums[i]+nums[j] == target {
				return []int{i, j}
			}

			j++
		}

		i++
	}

	return []int{}
}
*/

/*
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
*/

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
