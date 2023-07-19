package main

import "fmt"

func main() {
	fmt.Println("vim-go")
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	arr1 := []int{}
	arr2 := []int{}

	for l1 != nil {
		arr1 = append(arr1, l1.Val)
		l1 = l1.Next
	}

	for l2 != nil {
		arr2 = append(arr2, l2.Val)
		l2 = l2.Next
	}

	arr3 := []int{}
	n := 0

	idx1 := len(arr1) - 1
	idx2 := len(arr2) - 1
	next := &ListNode{}

	for idx1 >= 0 || idx2 >= 0 {
		sum := n
		if idx1 >= 0 {
			sum += arr1[idx1]
			idx1--
		}
		if idx2 >= 0 {
			sum += arr2[idx2]
			idx2--
		}

		if sum >= 10 {
			n = 1
		} else {
			n = 0
		}

		next.Val = sum % 10
		head := &ListNode{}
		head.Next = next
		next = head
		arr3 = append(arr3, sum%10)
	}

	if next.Val == 0 {
		return next.Next
	} else {
		return next
	}
}
