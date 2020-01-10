package main

import "fmt"

func main() {
	var n int

	n = 1230
	fmt.Println(isLucky(n))

	n = 239017
	fmt.Println(isLucky(n))

	n = 6446
	fmt.Println(isLucky(n))
}

func isLucky(n int) bool {
	num := 1
	count := -1

	clct := make([]int, 0)

	for num <= n {
		num = num * 10
		count++
	}

	mid := count / 2
	v1, v2 := 0, 0

	for count >= 0 {
		tenNum := pow(10, count)
		val := n / tenNum
		n = n - tenNum*val
		clct = append(clct, val)

		if count <= mid {
			v1 += val
		} else {
			v2 += val
		}

		count--
	}

	return (v1 == v2)
}

func pow(num int, times int) int {
	ret := 1

	for times > 0 {
		ret = ret * num
		times--
	}

	return ret
}
