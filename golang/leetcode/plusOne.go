// https://leetcode.com/problems/plus-one/submissions/

package main

import (
	"fmt"
)

func main() {

	var digits []int

	digits = []int{9, 9}
	fmt.Println(plusOne(digits))

	digits = []int{8, 9, 9, 9}
	fmt.Println(plusOne(digits))
}

func plusOne(digits []int) []int {
    digits[len(digits) - 1] += 1

    check(&digits, len(digits) - 1)
    
    return digits
}

func check(nums *[]int, idx int) {
    if (*nums)[idx] >= 10 {
        n := (*nums)[idx] % 10
        (*nums)[idx] = n
        
        if idx == 0 {
            (*nums) = append([]int{1}, (*nums)...)
        } else {
            (*nums)[idx - 1] += 1
            check(nums, idx - 1)
        }  
    }
}

/*
func plusOne(digits []int) []int {

	var ret []int

	i := len(digits) - 1
	digits[i] += 1

	plus := 0
	for i >= 0 {
		digits[i] += plus
		tmp := digits[i]
		fmt.Println(tmp, i)

		if tmp >= 10 {
			ret = append([]int{tmp % 10}, ret...)
			plus = tmp / 10
		} else {
			ret = append([]int{tmp}, ret...)
			plus = 0
		}

		i--
	}

	if plus > 0 {
		ret = append([]int{plus}, ret...)
	}

	return ret
}
*/
