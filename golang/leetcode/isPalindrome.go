package main

import (
	"fmt"
	"strings"
	"unicode"
)

func main() {
	// s := "aAzZ"
	// s := "aAczZ"
	// s := "aAcaa"
	s := "A man, a plan, a canal: Panama"
	// s := " "
	// s := "0P"
	// s := "a"

	// a 65
	// z 90
	// A 97
	// Z 122
	fmt.Println(isPalindrome(s))
}

func isPalindrome(s string) bool {
	s = strings.ToLower(s)
	start := 0
	end := len(s) - 1
	for start < end {
		for start < end && !unicode.IsLetter(rune(s[start])) && !unicode.IsNumber(rune(s[start])) {
			start++
		}
		for start < end && !unicode.IsLetter(rune(s[end])) && !unicode.IsNumber(rune(s[end])) {
			end--
		}

		if s[start] != s[end] {
			return false
		}

		start++
		end--
	}

	return true
}

/*
func isPalindrome(s string) bool {
	newS := ""

	for _, c := range s {
		// 0-9
		if c >= 48 && c <= 57 {
			newS = newS + string(c)
		}

		// uppercase
		if c >= 65 && c <= 90 {
			newS = newS + strings.ToLower(string(c))
		}

		// lowercase
		if c >= 97 && c <= 122 {
			newS = newS + string(c)
		}
	}

	start := 0
	end := len(newS) - 1
	for start < end {
		if newS[start] != newS[end] {
			return false
		}
		start = start + 1
		end = end - 1
	}

	return true
}
*/
