package main

import "fmt"

func main() {
	//fmt.Println(isPalindrome(1234))
	//fmt.Println(isPalindrome(1221))
	fmt.Println(isPalindrome(121))
	fmt.Println(isPalindrome(12321))
	//fmt.Println(isPalindrome(1212))
}

func isPalindrome(x int) bool {
	if x < 0 || (x != 0 && x%10 == 0) {
		return false
	}

	// check only Half of the digits of given Number to prevent from OVERFLOW
	check := 0
	for x > check {
		check = check*10 + x%10
		x /= 10
	}

	return (x == check) || (x == check/10)
}

/*
func isPalindrome(x int) bool {
	// if negative is false
	if x < 0 {
		return false
	}

	// Collect num in map to check is palindrome or not
	m := make(map[int]int)
	count := 0
	for x > 0 {
		digit := x % 10
		x /= 10

		m[count] = digit
		count++
	}

	start, end := 0, len(m)-1
	for start < end {
		if m[start] != m[end] {
			return false
		}

		start++
		end--
	}

	return true
}
*/
