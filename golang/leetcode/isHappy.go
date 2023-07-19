package main

import (
	"fmt"
)

func main() {
	fmt.Println(isHappy(1))
	fmt.Println(isHappy(2))
	fmt.Println(isHappy(19))
}

func isHappy(n int) bool {
	// Create a set to keep track of the numbers that have already been seen
	seen := make(map[int]bool)

	for n > 1 {
		if _, ok := seen[n]; ok {
			// n is seen again, so loops endlessly in a cycle which does not include 1
			return n == 1
		} else {
			// Add n to the set of seen numbers
			seen[n] = true
		}

		// Compute the sum of the squares of the digits of n
		sum := 0
		for n > 0 {
			digit := n % 10
			sum += digit * digit
			n /= 10
		}

		// Update n with the sum of the squares of its digits
		n = sum
	}

	// If n is 1, it is a happy number; otherwise, it is not
	return n == 1
}
