package main

import (
	"fmt"
)

type ListNode struct {
	Value interface{}
	Next  *ListNode
}

func main() {
	var ret, l *ListNode
	var n int

	l = &ListNode{Value: 1, Next: &ListNode{Value: 2, Next: &ListNode{Value: 3, Next: &ListNode{Value: 4, Next: &ListNode{Value: 5, Next: nil}}}}}
	n = 3

	ret = rearrangeLastN(l, n)
	fmt.Println("ans", ret)
	seeAll(ret)
}

/*
// other easier and faster way
func rearrangeLastN(l *ListNode, n int) *ListNode {
	if n == 0 {
		return l
	}
	var count int
	pointer1 := l
	pointer2 := l
	for pointer1.Next != nil {
		pointer1 = pointer1.Next
		count++
		if count >= n+1 {
			pointer2 = pointer2.Next
		}
	}

	if pointer2 == l && count < n {
		return l
	}

	head := pointer2.Next
	pointer2.Next = nil
	pointer1.Next = l
	return head
}
*/

func rearrangeLastN(l *ListNode, n int) *ListNode {
	if n == 0 || l == nil || l.Next == nil {
		return l
	}

	revs := revserse(l)
	curr := revs
	count := 0
	var part1, part2, prev *ListNode = nil, nil, nil

	part1 = revs

	for curr != nil {

		if count == n-1 {
			part2 = curr.Next
			fmt.Println("n", part1, part2, curr, prev)

			curr.Next = nil

			revs = revserse(part1)
			part1.Next = revserse(part2)
			break
		}

		prev = curr
		curr = curr.Next

		count++
	}

	return revs
}

func revserse(node *ListNode) *ListNode {
	var prev, next *ListNode = nil, nil
	curr := node

	for curr != nil {
		next = curr.Next
		curr.Next = prev
		prev = curr
		curr = next
	}

	return prev
}

func seeAll(node *ListNode) {

	for node != nil {
		fmt.Println(node.Value)

		node = node.Next
	}
}
