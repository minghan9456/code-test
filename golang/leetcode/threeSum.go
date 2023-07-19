package main

import (
	"fmt"
	"sort"
)

func main() {
	nums := []int{-1, 0, 1, 2, -1, -4}
	fmt.Println(threeSum(nums))
}

func threeSum(nums []int) [][]int {
	// Sort nums
	// Time: O(nlogn)
	sort.Slice(nums, func(a, b int) bool {
		return (nums[a] < nums[b])
	})

	// Time: O(n^2)
	// Enumerate nums[i]
	// Use 2 pointer to find all possible pairs of j, k such that
	// a. i < j < k
	// b. nums[i] + nums[j] + nums[k] == 0
	// target = 0 - nums[i]
	// if nums[j] + nums[k] > target, so k--
	// if nums[j] + nums[k] < target, so j++

	// Remove duplicates
	// Make sure nums[i] <= nums[j] <= nums[k]
	// Skip nums[i] if nums[i] == nums[i-1]

	n := len(nums)
	ans := [][]int{}

	for i := 0; i < n; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		j, k := i+1, n-1

		for j < k {
			target := 0 - nums[i]
			if nums[j]+nums[k] > target {
				k--
			} else if nums[j]+nums[k] < target {
				j++
			} else {
				ans = append(ans, []int{nums[i], nums[j], nums[k]})
				for j < k && nums[k] == nums[k-1] {
					k--
				}
				for j < k && nums[j] == nums[j+1] {
					j++
				}

				k--
				j++
			}
		}
	}

	return ans
}
