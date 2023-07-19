package main

import "fmt"

func main() {
	// fmt.Println(minFlips(2, 6, 5))
	fmt.Println(minFlips(26, 6, 5))
	// fmt.Println(minFlips(4, 2, 7))
	// fmt.Println(minFlips(1, 2, 3))
}

// TC: O(1)
// Iterate through each bit from 0 to 30 (as the problem states a,b,c <= 10^9 and 10^9 < 2^30) in numbers a, b and c. We then inspect the bit in c and consider each of the two cases listed above.
func minFlips(a int, b int, c int) int {
	flips := 0
	for i := 0; i <= 31; i++ {
		if (c>>i)&1 == 1 {
			// i'th bit of c is 1
			// i'th bit of a or i'th bit of b must be 1
			if (a>>i)&1 == 0 && (b>>i)&1 == 0 {
				flips++
			}
		} else {
			// i'th bit of c is 0
			// i'th bit of a and i'th bit of b must be 0
			if (a>>i)&1 == 1 {
				flips++
			}
			if (b>>i)&1 == 1 {
				flips++
			}
		}
	}

	return flips
}

/*
// TC: O(n)
func minFlips(a int, b int, c int) int {
	// a | b = c
	flips := 0
	i := 0
	for c>>i > 0 || a>>i > 0 || b>>i > 0 {
		if (c>>i)&1 == 1 {
			// i'th bit of c is 1
			// i'th bit of a OR i'th bit of b must be 1
			if (a>>i)&1 == 0 && (b>>i)&1 == 0 {
				fmt.Println("i", i)
				flips++
			}
		} else {
			// i'th bit of c is 0
			// i'th bit of a AND i'th bit of b must be 0

			if (a>>i)&1 == 1 {
				flips++
			}

			if (b>>i)&1 == 1 {
				flips++
			}
		}
		i++
	}
	return flips
}
*/
