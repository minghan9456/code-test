package main

import (
	"fmt"
	"sort"
)

type ListNode struct {
  Val int
  Next *ListNode
}

func main() {
  fmt.Println("******")

  l1 := &ListNode{Val : 4, Next : &ListNode{Val : 2, Next : &ListNode{Val : 1, Next : &ListNode{Val : 3, Next : nil}}}}
  showNodes(l1)

  fmt.Println("******")

  showNodes(sortList(l1))

  fmt.Println("******")
  showNodes(sortList(nil))
}

func sortList(head *ListNode) *ListNode {
  if head == nil {
    return nil
  }

  arr := make([]int, 0)
  for head != nil {
    arr = append(arr, head.Val)
    head = head.Next
  }

  sort.Ints(arr)
  arrLen := len(arr)

  newNode := &ListNode{}
  cur := newNode
  for i, v := range arr {
    cur.Val = v
    if i == arrLen - 1 {
      cur.Next = nil
    } else {
      cur.Next = &ListNode{}
    }

    cur = cur.Next
  }

  return newNode
}

func showNodes(node *ListNode) {
  if node != nil {
    target := node
    for {
      if (target.Next == nil) {
        fmt.Println(target)
       break
     } else {
       fmt.Println(target)
       target = target.Next
     }
    }
  }
}
