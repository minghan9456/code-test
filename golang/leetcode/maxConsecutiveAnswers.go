package main

import "fmt"

func main() {
	key := "TFFTFFFT"
	k := 1
	fmt.Println(maxConsecutiveAnswers(key, k))
}

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func min(a, b int) int {
	if a > b {
		return b
	} else {
		return a
	}
}

// Sliding Window
// TC : O(n)
// SC : O(1)
func maxConsecutiveAnswers(answerKey string, k int) int {
	// T 84, F 70

	tCount := 0
	fCount := 0
	left := 0
	right := 0
	maxL := 0

	for right < len(answerKey) {
		if answerKey[right] == 84 {
			tCount++
		} else {
			fCount++
		}

		right++

		if min(tCount, fCount) > k {
			if answerKey[left] == 84 {
				tCount--
			} else {
				fCount--
			}

			left++
		} else {
			maxL = max(maxL, tCount+fCount)
		}
	}

	return maxL
}
