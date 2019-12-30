package main

import "fmt"

func main() {
	fmt.Println(firstNotRepeatingCharacter("abacabad"))

	fmt.Println(firstNotRepeatingCharacter("abacabaabacaba"))
}

func firstNotRepeatingCharacter(s string) string {

	for i := 0; i <= len(s)-1; i++ {
		isMatch := false
		for j := 0; j <= len(s)-1; j++ {
			if i != j && s[i] == s[j] {
				isMatch = true
				break
			}
		}

		if !isMatch {
			return string(s[i])
		}
	}

	return "_"
}
