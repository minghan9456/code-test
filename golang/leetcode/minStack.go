package main

import (
	"fmt"
)

func main() {
	fmt.Println("vim-go")
}

type Elem struct {
	Val int
	Min int
}

type MinStack struct {
	List []Elem
}

func Constructor() MinStack {
	return MinStack{
		List: []Elem{},
	}
}

func (this *MinStack) Push(val int) {
	min := val
	if len(this.List) > 0 && this.GetMin() < val {
		min = this.GetMin()
	}

	e := Elem{
		Val: val,
		Min: min,
	}

	this.List = append(this.List, e)
}

func (this *MinStack) Pop() {
	this.List = this.List[:len(this.List)-1]
}

func (this *MinStack) Top() int {
	return this.List[len(this.List)-1].Val
}

func (this *MinStack) GetMin() int {
	return this.List[len(this.List)-1].Min
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(val);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */
