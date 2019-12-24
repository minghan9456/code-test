package main

import (
	"fmt"
	//"sort"
)

func main() {

	var nums []int

	nums = []int{1, 4, 3, 2}
	fmt.Println(arrayPairSum(nums))

	nums = []int{-5, -3, -4, -2, 1, 10}
	fmt.Println(arrayPairSum(nums))
}

func arrayPairSum(nums []int) int {
	var buckets [20001]int
	for _, v := range nums {
		buckets[v+10000]++
	}

	sum := 0
	odd := true
	for i := 0; i < len(buckets); i++ {
		for buckets[i] > 0 {
			if odd {
				fmt.Println(i)
				sum += i - 10000
			}
			odd = !odd
			buckets[i]--
		}
	}
	return sum
}

/*
func arrayPairSum(nums []int) int {
	sum := 0

	sort.Ints(nums)

	for i := 0; i < len(nums); i += 2 {
		if nums[i] < nums[i+1] {
			sum += nums[i]
		} else {
			sum += nums[i+1]
		}
	}

	return sum
}
*/
