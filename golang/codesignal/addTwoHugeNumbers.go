package main

import "fmt"

type ListNode struct {
	Value interface{}
	Next  *ListNode
}

func main() {
	var a, b, ret *ListNode

	a = &ListNode{Value: 9876, Next: &ListNode{Value: 5432, Next: &ListNode{Value: 1999, Next: nil}}}
	b = &ListNode{Value: 1, Next: &ListNode{Value: 8001, Next: nil}}

	ret = addTwoHugeNumbers(a, b)
	fmt.Println(ret)
	seeAll(ret)

	a = &ListNode{Value: 123, Next: &ListNode{Value: 4, Next: &ListNode{Value: 5, Next: nil}}}
	b = &ListNode{Value: 100, Next: &ListNode{Value: 100, Next: &ListNode{Value: 100, Next: nil}}}

	ret = addTwoHugeNumbers(a, b)
	fmt.Println(ret)
	seeAll(ret)

	/*
		a = &ListNode{Value: 1234, Next: &ListNode{Value: 1234, Next: &ListNode{Value: 0, Next: nil}}}
		b = &ListNode{Value: 0, Next: nil}

		ret = addTwoHugeNumbers(a, b)
		fmt.Println(ret)
		seeAll(ret)

		a = &ListNode{Value: 0, Next: nil}
		b = &ListNode{Value: 0, Next: nil}

		ret = addTwoHugeNumbers(a, b)
		fmt.Println(ret)
		seeAll(ret)
	*/

	a = &ListNode{Value: 2, Next: nil}
	b = &ListNode{Value: 1, Next: &ListNode{Value: 9999, Next: nil}}

	ret = addTwoHugeNumbers(a, b)
	fmt.Println(ret)
	seeAll(ret)

}

func addTwoHugeNumbers(a *ListNode, b *ListNode) *ListNode {
	a = revserse(a)
	b = revserse(b)

	var ret *ListNode = &ListNode{Value: 0, Next: nil}
	curr := ret

	for a != nil || b != nil {
		curr.Next = &ListNode{Value: 0, Next: nil}

		if a != nil {
			//fmt.Println("a", a.Value)

			curr.Value = curr.Value.(int) + a.Value.(int)
			a = a.Next
		}

		if b != nil {
			//fmt.Println("b", b.Value)

			curr.Value = curr.Value.(int) + b.Value.(int)
			b = b.Next
		}

		if curr.Value.(int) >= 10000 {
			curr.Next.Value = curr.Value.(int) / 10000
			curr.Value = curr.Value.(int) % 10000
		}

		//fmt.Println("curr", curr.Value)

		curr = curr.Next
	}

	ret = revserse(ret)
	if ret.Value.(int) == 0 {
		return ret.Next
	}

	return ret
}

func seeAll(node *ListNode) {

	for node != nil {
		fmt.Println(node.Value)

		node = node.Next
	}
}

func revserse(node *ListNode) *ListNode {
	var prev *ListNode = nil
	curr := node
	var next *ListNode = nil

	for curr != nil {
		next = curr.Next

		curr.Next = prev

		prev = curr

		curr = next
	}

	return prev
}
