package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println("vim-go")
}

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func getMinimumDifference(root *TreeNode) int {
	ans := math.MaxInt32

	if root == nil {
		return 0
	}

	vals := []int{}
	queue := []*TreeNode{}
	curr := root

	for curr != nil || len(queue) > 0 {
		if curr != nil {
			queue = append(queue, curr)
			curr = curr.Left
		} else {
			node := queue[len(queue)-1]
			vals = append(vals, node.Val)
			queue = queue[:len(queue)-1]
			curr = node.Right
		}
	}

	ans = vals[1] - vals[0]
	for i := 1; i < len(vals)-1; i++ {
		if vals[i+1]-vals[i] < ans {
			ans = vals[i+1] - vals[i]
		}
	}

	return ans
}

func getMinimumDifference(root *TreeNode) int {
  ans := math.MaxInt32

  if root == nil {
    return 0
  }

  prev := -1
  queue := []*TreeNode{}
  curr := root

  for curr != nil || len(queue) > 0 {
    if curr != nil {
      queue = append(queue, curr)
      curr = curr.Left
    } else {
      node := queue[len(queue)-1]

      if prev != -1 && node.Val - prev < ans{
        ans = node.Val - prev
      }
      prev = node.Val

      queue = queue[:len(queue)-1]
      curr = node.Right
    }
  }

  return ans
}
