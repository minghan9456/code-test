package main

import "fmt"

func main() {
	nums := []int{7, 4, 3, 9, 1, 8, 5, 2, 6}
	k := 3
	// nums := []int{7}
	// k := 0
	fmt.Println(getAverages(nums, k))
}

/*
TC: O(n)
SC: O(n)
*/
func getAverages(nums []int, k int) []int {
	n := len(nums)
	length := k*2 + 1
	ans := make([]int, n)
	sum := 0

	for i := 0; i < n; i++ {
		ans[i] = -1
	}

	if n < length {
		return ans
	}

	left := 0
	for right := 0; right < n; right++ {
		sum += nums[right]

		if right-left+1 >= length {
			ans[left+k] = sum / length
			sum -= nums[left]
			left++
		}

	}

	return ans
}

/*
TC: O(n*k)
SC: O(n)
*/
/*
func getAverages(nums []int, k int) []int {
	ans := []int{}
	n := len(nums)

	for i := 0; i < n; i++ {

		if i < k || i >= n-k {
			ans = append(ans, -1)
			continue
		}

		sum := nums[i]
		for j := 1; j <= k; j++ {
			sum += nums[i-j]
			sum += nums[i+j]
		}

		ans = append(ans, sum/(k*2+1))
	}

	return ans
}
*/
