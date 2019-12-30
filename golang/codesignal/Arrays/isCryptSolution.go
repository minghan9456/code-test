package main

import "fmt"

func main() {

	var crypt []string
	var solution [][]string

	crypt = []string{"SEND", "MORE", "MONEY"}
	solution = [][]string{
		{"O", "0"},
		{"M", "1"},
		{"Y", "2"},
		{"E", "5"},
		{"N", "6"},
		{"D", "7"},
		{"R", "8"},
		{"S", "9"},
	}

	fmt.Println(isCryptSolution(crypt, solution))

	crypt = []string{"TEN", "TWO", "ONE"}
	solution = [][]string{
		{"O", "1"},
		{"T", "0"},
		{"W", "9"},
		{"E", "5"},
		{"N", "4"},
	}

	fmt.Println(isCryptSolution(crypt, solution))
}

func isCryptSolution(crypt []string, solution [][]string) bool {
	newSolution := make(map[string]string)

	for i := 0; i <= len(solution)-1; i++ {
		newSolution[solution[i][0]] = solution[i][1]
	}

	nums := make([]int, 0)

	for i := 0; i <= len(crypt)-1; i++ {
		num := 0

		for j := 0; j <= len(crypt[i])-1; j++ {
			char := string(crypt[i][j])

			val, ok := newSolution[char]
			if ok {
				if j == 0 && val == "0" && len(crypt[i]) != 1 {
					return false
				}
			} else {
				return false
			}

			tmp := int(val[0] - "0"[0])
			ten := Pow(10, len(crypt[i])-1-j)
			num = num + (tmp)*ten
		}

		nums = append(nums, num)
	}

	return (nums[0]+nums[1] == nums[2])
}

func Pow(x, y int) int {
	ret := 1

	for i := y - 1; i >= 0; i-- {
		ret = ret * x
	}

	return ret
}
