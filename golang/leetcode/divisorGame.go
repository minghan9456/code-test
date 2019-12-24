package main

import (
	"fmt"
)

func main() {

	var N int

	N = 3
	fmt.Println(divisorGame(N))
}

func divisorGame(N int) bool {
	return N%2 == 0
}
