package main

import (
	"fmt"
)

/*
	What if the given array is already sorted? How would you optimize your algorithm?
	What if nums1's size is small compared to nums2's size? Which algorithm is better?
	What if elements of nums2 are stored on disk, and the memory is limited such that you cannot load all elements into the memory at once?
*/

func main() {
	var nums1, nums2 []int

	nums1 = []int{4, 9, 5}
	nums2 = []int{9, 4, 9, 8, 4}
	fmt.Println(intersect(nums1, nums2))

	nums1 = []int{1}
	nums2 = []int{1}
	fmt.Println(intersect(nums1, nums2))
}

func intersect(nums1 []int, nums2 []int) []int {
	var ret []int

	if len(nums1) < 1 || len(nums2) < 1 {
		return ret
	}

	j := 0
	for j <= len(nums1)-1 {

		i := 0
		len2 := len(nums2) - 1
		for i <= len2 && len2 >= 0 {
			if nums1[j] == nums2[i] {
				ret = append(ret, nums1[j])
				nums2 = append(nums2[0:i], nums2[i+1:]...)
				break
			} else {
				i++
			}
		}

		j++
	}

	return ret
}
