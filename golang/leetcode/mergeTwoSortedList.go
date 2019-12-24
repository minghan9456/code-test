package main

import (
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	/*
		ll := &ListNode{Val : 1, Next : &ListNode{Val : 2, Next : &ListNode{Val : 4, Next : nil}}}
		ll = bePrev(10, ll)
		showNodes(ll)
	*/
	fmt.Println("******")

	l1 := &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 4, Next: nil}}}
	l2 := &ListNode{Val: 1, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4, Next: nil}}}
	showNodes(mergeTwoLists(l1, l2))

	fmt.Println("******")

	l3 := &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 4, Next: nil}}}
	l4 := &ListNode{Val: 3, Next: &ListNode{Val: 5, Next: &ListNode{Val: 6, Next: nil}}}
	showNodes(mergeTwoLists(l3, l4))
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	res := &ListNode{}
	cur := res

	for l1 != nil || l2 != nil {
		if l1 == nil {
			cur.Next = l2
			l2 = nil
			continue
		}

		if l2 == nil {
			cur.Next = l1
			l1 = nil
			continue
		}

		if l1.Val > l2.Val {
			cur.Next = l2
			l2 = l2.Next
		} else {
			cur.Next = l1
			l1 = l1.Next
		}
		cur = cur.Next
	}

	return res.Next
}

func showNodes(node *ListNode) {
	target := node
	for {
		if target.Next == nil {
			fmt.Println(target)
			break
		} else {
			fmt.Println(target)
			target = target.Next
		}
	}
}
