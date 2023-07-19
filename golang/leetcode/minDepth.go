package main

import "fmt"

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

type Elem struct {
	Layer int
	Node  *TreeNode
}

// BFS
func minDepth(root *TreeNode) int {
	queue := []*Elem{}
	elem := &Elem{Layer: 1, Node: root}
	queue = append(queue, elem)

	for len(queue) > 0 {
		curr := queue[0]
		// fmt.Println(curr.Layer, curr.Node.Val)
		queue = queue[1:]

		if curr.Node == nil {
			continue
		}

		// is leaf.
		if curr.Node.Left == nil && curr.Node.Right == nil {
			return curr.Layer
		}

		if curr.Node.Left != nil {
			queue = append(queue, &Elem{Layer: curr.Layer + 1, Node: curr.Node.Left})
		}
		if curr.Node.Right != nil {
			queue = append(queue, &Elem{Layer: curr.Layer + 1, Node: curr.Node.Right})
		}
	}

	return 0
}
