package main

import "fmt"

func main() {
	// fmt.Println(maxValue(6, 1, 10))
	// fmt.Println(maxValue(8, 2, 30))
	fmt.Println(maxValue(1, 0, 24))
}

// sum of numbers from [1 to n]
func calc(n int) int {
	return ((n * (n + 1)) / 2)
}

func getSum(totalLen, val int) int {
	totalSum := 0

	decreasingLen := val

	if totalLen >= decreasingLen {
		//[val, val - 1, val - 2, ... 1, 1, 1, 1]
		currSum := calc(val)
		remainingSum := totalLen - decreasingLen //all 1's
		totalSum = currSum + remainingSum
	} else {
		//[val, val - 1, val - 2]
		lastVal := decreasingLen - totalLen
		totalSum = calc(val) - calc(lastVal)
	}
	return totalSum
}

func maxValue(n int, index int, maxSum int) int {
	ans := -1
	left := 1
	right := maxSum
	leftCount := index
	rightCount := n - index - 1
	fmt.Println(leftCount, rightCount)

	for left <= right {
		mid := (left + right + 1) / 2

		leftSum := getSum(leftCount, mid-1)
		rightSum := getSum(rightCount, mid-1)
		fmt.Println(leftSum, mid, rightSum)
		totalSum := leftSum + mid + rightSum

		if totalSum > maxSum {
			right = mid - 1
		} else {
			ans = mid
			left = mid + 1
		}

	}

	return ans
}
