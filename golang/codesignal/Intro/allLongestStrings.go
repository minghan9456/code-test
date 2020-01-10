package main

import "fmt"

func main() {
	var inputArray []string
	inputArray = []string{
		"aa",
		"aba",
		"aa",
		"ad",
		"vcd",
		"aba",
	}

	fmt.Println(allLongestStrings(inputArray))
}

func allLongestStrings(inputArray []string) []string {
	ret := make([]string, 0)
	maxCount := 0
	for _, val := range inputArray {
		if len(val) > maxCount {
			// reset ret
			ret = []string{val}
			maxCount = len(val)
		} else if len(val) == maxCount {
			ret = append(ret, val)
		}
	}
	return ret
}

/*
func allLongestStrings(inputArray []string) []string {

	var max int
	var ret []string

	for i := 0; i <= len(inputArray)-1; i++ {
		if len(inputArray[i]) > max {
			max = len(inputArray[i])
		}
	}

	for i := 0; i <= len(inputArray)-1; i++ {
		if len(inputArray[i]) >= max {
			ret = append(ret, inputArray[i])
		}
	}

	return ret
}
*/
