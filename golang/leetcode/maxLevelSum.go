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
func maxLevelSum(root *TreeNode) int {
	queue := []*TreeNode{root}
	level := 0
	maxLevel := 0
	max := math.MinInt32

	for len(queue) > 0 {
		level += 1
		sum := 0

		n := len(queue)
		for i := 0; i < n; i++ {
			curr := queue[0]
			queue = queue[1:]

			sum += curr.Val

			if curr.Left != nil {
				queue = append(queue, curr.Left)
			}
			if curr.Right != nil {
				queue = append(queue, curr.Right)
			}
		}

		if sum > max {
			max = sum
			maxLevel = level
		}
	}

	return maxLevel
}

type Elem struct {
	Node  *TreeNode
	Level int
}

func maxLevelSum(root *treenode) int {
	queue := []Elem{Elem{Node: root, Level: 1}}

	ans := []int{}

	for len(queue) > 0 {
		curr := queue[0]
		queue = queue[1:]

		if curr.Node != nil {
			fmt.Println(curr.Node.Val, curr.Level)
			if len(ans) < curr.Level {
				ans = append(ans, 0)
			}
			ans[curr.Level-1] += curr.Node.Val

			if curr.Node.Left != nil {
				queue = append(queue, Elem{Node: curr.Node.Left, Level: curr.Level + 1})
			}
			if curr.Node.Right != nil {
				queue = append(queue, Elem{Node: curr.Node.Right, Level: curr.Level + 1})
			}

		}
	}

	fmt.Println(ans)

	max := math.MinInt32
	pos := 1
	for i := 0; i < len(ans); i++ {
		if ans[i] > max {
			max = ans[i]
			pos = i + 1
		}
	}

	return pos
}
