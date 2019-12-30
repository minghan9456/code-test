package main

import (
	"fmt"
	"os"
)

type ListNode struct {
	Value interface{}
	Next  *ListNode
}

func main() {
	var ret, l *ListNode
	var k int

	/*
		l = &ListNode{Value: 1, Next: &ListNode{Value: 2, Next: &ListNode{Value: 3, Next: &ListNode{Value: 4, Next: &ListNode{Value: 5, Next: &ListNode{Value: 6, Next: nil}}}}}}
		k = 2
	*/

	l = &ListNode{Value: 1, Next: &ListNode{Value: 2, Next: &ListNode{Value: 3, Next: &ListNode{Value: 4, Next: &ListNode{Value: 5, Next: nil}}}}}
	k = 3

	/*
		l = &ListNode{Value: 1, Next: &ListNode{Value: 2, Next: &ListNode{Value: 3, Next: &ListNode{Value: 4, Next: &ListNode{Value: 5, Next: nil}}}}}
		k = 6
	*/

	/*
		l = &ListNode{Value: 1, Next: nil}
		k = 1
	*/

	ret = reverseNodesInKGroups(l, k)
	fmt.Println("ans", ret)
	seeAll(ret)

	os.Exit(2)
}

func reverseNodesInKGroups(l *ListNode, k int) *ListNode {
	curr := l
	var subHead, preHead, ret *ListNode = nil, nil, &ListNode{Value: 0, Next: nil}
	preHead = ret
	isFullk, isEnd := false, false

	if k == 1 || l == nil || l.Next == nil {
		return l
	}

	for !isEnd {
		subHead, isEnd, isFullk = takeKNodes(curr, k)
		fmt.Println(subHead, isEnd, isFullk)

		if isFullk {
			preHead.Next = revserse(subHead, k)
			//fmt.Println(preHead)
			//seeAll(preHead)
			curr = subHead.Next
			preHead = subHead
		}
	}

	if ret.Next == nil {
		return l
	} else {
		return ret.Next
	}
}

func revserse(node *ListNode, k int) *ListNode {
	var prev *ListNode = nil
	curr := node
	var next *ListNode = nil
	count := 0

	for curr != nil {
		next, curr.Next = curr.Next, prev
		prev, curr = curr, next
		if count == k-1 {
			break
		} else {
			count++
		}
	}
	node.Next = curr

	return prev
}

func takeKNodes(head *ListNode, k int) (*ListNode, bool, bool) {
	curr := head
	count := 0

	for curr != nil && count < k {
		curr, count = curr.Next, count+1
	}

	return head, (curr == nil), (count%k == 0)
}

func seeAll(node *ListNode) {

	for node != nil {
		fmt.Println(node.Value)

		node = node.Next
	}
}
