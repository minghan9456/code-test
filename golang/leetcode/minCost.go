package main

import "fmt"

func main() {
	cuts := []int{3, 5, 1, 4}
	// cuts := []int{3, 5}
	fmt.Println(minCost(7, cuts))
}

func minCost(n int, cuts []int) int {
	cost := 0

	r(0, n, 0, cuts, &cost)

	return cost
}

func r(start, end, index int, cuts []int, tcost *int) int {
	cost := 0

	if index <= len(cuts)-1 {
		cut := cuts[index]

		if start < cut && cut < end {
			cost = end - start
			*tcost += cost
			fmt.Println("cost", cost)
			index++

			//fmt.Println(cut, end, index, "right")
			// fmt.Println(start, cut, index, "left")
			if index <= len(cuts)-1 {
				return r(cut, end, index, cuts, tcost) + r(start, cut, index, cuts, tcost)
			} else {
				return cost
			}
		} else {
			// fmt.Println(start, end, cut, index, "no")
			index++
			return r(start, end, index, cuts, tcost)
		}
	} else {
		return cost
	}
}
