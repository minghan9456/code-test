package main

import (
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {

	var head *ListNode

	head = &ListNode{Val: 0, Next: &ListNode{Val: 1, Next: &ListNode{Val: 1, Next: nil}}}

	fmt.Println(getDecimalValue(head))
}

func getDecimalValue(head *ListNode) int {
	var ret int = 0

	for head != nil {
		tmp := ret << 1
		fmt.Println(tmp, head.Val)
		ret = tmp | head.Val
		fmt.Println("ret", ret)
		head = head.Next
	}

	return ret
}

/*
func getDecimalValue(head *ListNode) int {
	var ret, count, num int = 0, 0, 1
	var curr *ListNode = nil

	curr = head
	for curr.Next != nil {
		curr = curr.Next
		count++
		num = num << 1
	}

	curr = head
	for num >= 1 && curr != nil {
		ret = ret + num*curr.Val
		curr = curr.Next
		num = num >> 1
	}

	return ret
}
*/

func seeAll(node *ListNode) {

	for node != nil {
		fmt.Println(node.Val)

		node = node.Next
	}
}
