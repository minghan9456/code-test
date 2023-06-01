package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	// s := "-42"
	// s := "  -43"
	// s := "4193 with words"
	// s := "words and 987"
	// s := "34+3"
	// s := "a23fd"
	// s := "  0000000000012345678"
	// s := "050678"
	s := "2147483648"

	fmt.Println(strconv.Atoi(s))
	fmt.Println(myAtoi(s))
}

func myAtoi(s string) int {
	const MaxUint = ^uint32(0)
	const MaxInt = int(MaxUint >> 1)
	const MinInt = -MaxInt - 1
	plus := true
	s = strings.TrimSpace(s)

	if len(s) == 0 {
		return 0
	}

	if s[0] == '-' {
		plus = false
		s = s[1:]
	} else if s[0] == '+' {
		plus = true
		s = s[1:]
	}

	start := false
	nStr := ""
	for _, ch := range []byte(s) {
		r := ch - '0'
		if !start && r > 0 {
			start = true
		}

		if start {
			if r <= 9 {
				nStr += string(ch)
			} else {
				break
			}
		}
	}

	n := 0
	for _, ch := range []byte(nStr) {
		ch -= '0'
		if ch <= 9 {
			n = n*10 + int(ch)
			if n > MaxInt {
				if plus {
					return MaxInt
				} else {
					return MinInt
				}
			}
		}
	}

	if !plus {
		n = -n
	}
	return n

}
