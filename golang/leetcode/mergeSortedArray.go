package main

import "fmt"

func main() {
	/*
		nums1 := []int{1, 2, 19, 20, 0, 0, 0}
		m := 4
		nums2 := []int{4, 5, 6}
		n := 3
	*/

	nums1 := []int{1, 2, 19, 20, 0}
	m := 4
	nums2 := []int{4}
	n := 1
	merge(nums1, m, nums2, n)
	fmt.Println(nums1)
}

func merge(nums1 []int, m int, nums2 []int, n int) {

	x := m - 1
	y := n - 1
	z := m + n - 1

	for x >= 0 && y >= 0 {
		if nums1[x] > nums2[y] {
			nums1[z] = nums1[x]
			x--
		} else {
			nums1[z] = nums2[y]
			y--
		}

		z--
	}

	if y >= 0 {
		for i := 0; i <= y; i++ {
			nums1[i] = nums2[i]
		}
	}
}
