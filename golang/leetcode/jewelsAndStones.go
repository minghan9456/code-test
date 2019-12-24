package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(numJewelsInStones("aA", "aAAbbbb"))
	fmt.Println(numJewelsInStones("z", "ZZ"))
	fmt.Println(numJewelsInStones("Abs", "aaBefsAbs"))
}

func numJewelsInStones(J string, S string) int {

	num := 0
	beforeLen := 0
	afeterLen := 0

	for _, word := range J {
		beforeLen = len(S)
		S = strings.Replace(S, string(word), "", -1)
		afeterLen = len(S)
		if beforeLen > afeterLen {
			num += beforeLen - afeterLen
		}
	}
	return num
}
