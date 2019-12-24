package main

import (
	"fmt"
	"sort"
)

func main() {

	var num int = 553
	var nums []int
	nums = make([]int, 0)

	nums = IntToSlice(num, nums)

	//nums := []int{1, 2, 3}
	//nums := []int{3, 5, 6, 7, 1}
	//nums := []int{3, 2, 1}
	//nums := []int{1, 2, 6, 4, 2}
	//nums := []int{1, 2, 0, 3, 0, 1, 2, 4}
	fmt.Println(nums)
	nextPermutation(nums)
	fmt.Println(nums)
	SliceToInt(nums)
}

func SliceToInt(sequence []int) int {
	sum := 0
	length := len(sequence) - 1
	tenPow := 1

	for i := length; i >= 0; i-- {
		sum += tenPow * sequence[i]
		tenPow = tenPow * 10
	}

	return sum
}

func IntToSlice(n int, sequence []int) []int {
	if n != 0 {
		i := n % 10
		// sequence = append(sequence, i) // reverse order output
		sequence = append([]int{i}, sequence...)
		return IntToSlice(n/10, sequence)
	}
	return sequence
}

func nextPermutation(nums []int) {
	length := len(nums)
	if length == 0 {
		return
	}

	targetIdx := length - 2
	for targetIdx >= 0 {
		if nums[targetIdx] < nums[targetIdx+1] {
			break
		}
		targetIdx--
	}

	if targetIdx < 0 {
		//sort.Ints(nums)
		return
	}

	lgrIdx := targetIdx + 1
	for lgrIdx < length && nums[lgrIdx] > nums[targetIdx] {
		lgrIdx++
	}

	tmpInt := nums[targetIdx]
	nums[targetIdx] = nums[lgrIdx-1]
	nums[lgrIdx-1] = tmpInt

	sort.Ints(nums[targetIdx+1 : length])
}
