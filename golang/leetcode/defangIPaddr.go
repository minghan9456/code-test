package main

import (
	"fmt"
)

func main() {
	var address string

	address = "1.2.1.1"
	fmt.Println(defangIPaddr(address))

	address = "255.100.50.0"
	fmt.Println(defangIPaddr(address))
}

func defangIPaddr(address string) string {
	retStr := ""

	for _, charRune := range address {
		if charRune == 46 {
			retStr += "[.]"
		} else {
			retStr += string(charRune)
		}
	}

	return retStr
}
