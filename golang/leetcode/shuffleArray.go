package main

import (
	"math/rand"
	"time"
)

func main() {
}

type Solution struct {
	Origin []int
}

func Constructor(nums []int) Solution {
	return Solution{nums}
}

func (this *Solution) Reset() []int {
	return this.Origin
}

func (this *Solution) Shuffle() []int {
	n := len(this.Origin)
	dest := make([]int, n)

	rand.Seed(time.Now().UnixNano())
	perm := rand.Perm(n)
	for i, v := range perm {
		dest[v] = src[i]
	}

	// fmt.Println(src)
	// fmt.Println(dest)
	return dest
}

/**
 * Your Solution object will be instantiated and called as such:
 * obj := Constructor(nums);
 * param_1 := obj.Reset();
 * param_2 := obj.Shuffle();
 */
