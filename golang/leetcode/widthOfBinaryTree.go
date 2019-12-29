package main

import (
	"fmt"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {

	var root *TreeNode
	var ret int

	/*
		root = &TreeNode{
			Val: 1,
			Left: &TreeNode{
				Val: 3,
				Left: &TreeNode{
					Val:   5,
					Left:  nil,
					Right: nil,
				},
				Right: nil,
			},
			Right: &TreeNode{
				Val:   2,
				Left:  nil,
				Right: nil,
			},
		}

		ret = widthOfBinaryTree(root)
		fmt.Println(ret)
	*/

	root = &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 3,
			Left: &TreeNode{
				Val: 5,
				Left: &TreeNode{
					Val:   6,
					Left:  nil,
					Right: nil,
				},
				Right: nil,
			},
			Right: nil,
		},
		Right: &TreeNode{
			Val:  2,
			Left: nil,
			Right: &TreeNode{
				Val:  9,
				Left: nil,
				Right: &TreeNode{
					Val:   7,
					Left:  nil,
					Right: nil,
				},
			},
		},
	}

	ret = widthOfBinaryTree(root)
	fmt.Println(ret)
}

func widthOfBinaryTree(root *TreeNode) int {
	clct := make(map[int][]int)

	printPreorder(root, 0, 0, clct)
	fmt.Println(clct)

	width := 1

	for _, arr := range clct {
		if len(arr) >= 2 {
			if width < arr[len(arr)-1]-arr[0]+1 {
				width = arr[len(arr)-1] - arr[0] + 1
			}
		}
	}

	return width
}

func printPreorder(node *TreeNode, level int, pos int, clct map[int][]int) {

	if node == nil {
		return
	}

	fmt.Println(node.Val, level, pos)

	_, ok := clct[level]
	if ok {
		clct[level] = append(clct[level], pos)
	} else {
		clct[level] = []int{pos}
	}

	printPreorder(node.Left, level+1, pos*2, clct)

	printPreorder(node.Right, level+1, pos*2+1, clct)
}
