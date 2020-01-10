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
