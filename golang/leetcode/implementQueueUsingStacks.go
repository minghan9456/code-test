package main

import (
	// "fmt"
)

func main() {
   /* 
   obj := Constructor();
   obj.Push(x);
   param_2 := obj.Pop();
   param_3 := obj.Peek();
   param_4 := obj.Empty();
   */
}

type Stack []int
func (s *Stack) Empty() bool {
    return len(*s) == 0
}
func (s *Stack) Push(x int) {
    *s = append(*s, x)
} 
func (s *Stack) Pop() int {
    last := (*s)[len(*s) - 1]
    *s = (*s)[:len(*s) - 1]
    return last
}
func (s *Stack) Peek() int {
    return (*s)[len(*s) - 1]
}

type MyQueue struct {
    s1, s2 Stack
}


/** Initialize your data structure here. */
func Constructor() MyQueue {
    return MyQueue{}
}


/** Push element x to the back of queue. */
func (this *MyQueue) Push(x int)  {
    this.s1.Push(x)
}


/** Removes the element from in front of queue and returns that element. */
func (this *MyQueue) Pop() int {
    if this.s2.Empty() {
        for !this.s1.Empty() {
            this.s2.Push(this.s1.Pop())
        }
    }
    
    return this.s2.Pop()
}


/** Get the front element. */
func (this *MyQueue) Peek() int {
    if this.s2.Empty() {
        for !this.s1.Empty() {
            this.s2.Push(this.s1.Pop())
        }
    }
    
    return this.s2.Peek()
}


/** Returns whether the queue is empty. */
func (this *MyQueue) Empty() bool {
    return this.s1.Empty() && this.s2.Empty()
}


/**
 * Your MyQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Peek();
 * param_4 := obj.Empty();
 */
