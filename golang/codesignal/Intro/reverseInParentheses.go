package main

import (
	"fmt"
)

func main() {
	var inputString string

	inputString = "foo(bar(baz))blim"
	//inputString = "(abc)d(efg)"
	fmt.Println("ans", reverseInParentheses(inputString))
}

func reverseInParentheses(inputString string) string {

	if len(inputString) == 0 {
		return inputString
	}

	arr := []int{}
	i := 0

	for i <= len(inputString)-1 {
		if inputString[i] == 40 {
			arr = append(arr, i)
		}

		if inputString[i] == 41 {
			left := arr[len(arr)-1]
			//fmt.Println(inputString[left+1 : i])
			reverseStr := reverse(inputString[left+1 : i])
			//fmt.Println(left, i, reverseStr)
			//fmt.Println(inputString[0:left] + reverseStr + inputString[i+1:])

			return reverseInParentheses(inputString[0:left] + reverseStr + inputString[i+1:])
		}

		i++
	}

	return inputString
}

func reverse(str string) string {
	reverseStr := ""

	for i := len(str) - 1; i >= 0; i-- {
		if str[i] == 40 {
			reverseStr += string(41)
		} else if str[i] == 41 {
			reverseStr += string(40)
		} else {
			reverseStr += string(str[i])
		}

	}

	return reverseStr
}
