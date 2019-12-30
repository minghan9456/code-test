package main

import (
	"fmt"
)

type ListNode struct {
	Value interface{}
	Next  *ListNode
}

func main() {
	var ret, l1, l2 *ListNode

	/*
		l1 = &ListNode{Value: 1, Next: &ListNode{Value: 2, Next: &ListNode{Value: 3, Next: nil}}}
		l2 = &ListNode{Value: 4, Next: &ListNode{Value: 5, Next: &ListNode{Value: 6, Next: nil}}}

		ret = mergeTwoLinkedLists(l1, l2)
		fmt.Println(ret)
		seeAll(ret)
	*/

	/*
		l1 = &ListNode{Value: 1, Next: &ListNode{Value: 3, Next: &ListNode{Value: 5, Next: nil}}}
		l2 = &ListNode{Value: 0, Next: &ListNode{Value: 2, Next: &ListNode{Value: 4, Next: nil}}}

		ret = mergeTwoLinkedLists(l1, l2)
		fmt.Println(ret)
		seeAll(ret)
	*/

	/*
		l1 = &ListNode{Value: 1, Next: &ListNode{Value: 1, Next: &ListNode{Value: 2, Next: &ListNode{Value: 4, Next: nil}}}}
		l2 = &ListNode{Value: 0, Next: &ListNode{Value: 3, Next: &ListNode{Value: 5, Next: nil}}}

		ret = mergeTwoLinkedLists(l1, l2)
		fmt.Println(ret)
		seeAll(ret)
	*/

	/*
		//l1 = &ListNode{Value: 1, Next: &ListNode{Value: 1, Next: &ListNode{Value: 2, Next: &ListNode{Value: 4, Next: nil}}}}
		l2 = &ListNode{Value: 0, Next: &ListNode{Value: 3, Next: &ListNode{Value: 5, Next: nil}}}

		ret = mergeTwoLinkedLists(l1, l2)
		fmt.Println(ret)
		seeAll(ret)
	*/

	l1 = &ListNode{Value: -1, Next: &ListNode{Value: -1, Next: &ListNode{Value: 0, Next: &ListNode{Value: 1, Next: nil}}}}
	l2 = &ListNode{Value: -1, Next: &ListNode{Value: 0, Next: &ListNode{Value: 0, Next: &ListNode{Value: 1, Next: &ListNode{Value: 1, Next: nil}}}}}

	ret = mergeTwoLinkedLists(l1, l2)
	fmt.Println(ret)
	seeAll(ret)
}

func mergeTwoLinkedLists(l1 *ListNode, l2 *ListNode) *ListNode {
	var prev *ListNode = nil
	curr := l1

	for l2 != nil {

		if curr != nil {
			if curr.Value.(int) > l2.Value.(int) {
				newNode := &ListNode{Value: l2.Value, Next: nil}
				if prev == nil {
					prev = newNode
					prev.Next = curr
					l1 = prev
				} else {
					prev.Next = newNode
					prev = prev.Next
					prev.Next = curr
				}

				l2 = l2.Next
			} else {
				prev = curr
				curr = curr.Next
			}
		} else {
			newNode := &ListNode{Value: l2.Value, Next: nil}

			if prev == nil {
				prev = newNode
				l1 = prev
			} else {
				prev.Next = newNode
			}
			curr = newNode

			l2 = l2.Next
		}

	}

	return l1
}

func seeAll(node *ListNode) {

	for node != nil {
		fmt.Println(node.Value)

		node = node.Next
	}
}
