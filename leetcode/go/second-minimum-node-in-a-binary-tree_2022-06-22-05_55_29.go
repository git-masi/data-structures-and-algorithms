// https://leetcode.com/problems/second-minimum-node-in-a-binary-tree/

// Overall this question was pretty easy though I missed one condition
// the first time through:
// 		node.Val != *smallest && node.Val != *second
// Once I spotted that it was a simple correction and then everything
// worked as expected

package main

import (
	"fmt"
	"math"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	x := TreeNode{
		Val: 2,
		Left: &TreeNode{
			Val: 2,
		},
		Right: &TreeNode{
			Val: 5,
			Left: &TreeNode{
				Val: 5,
			},
			Right: &TreeNode{
				Val: 7,
			},
		},
	}

	fmt.Println(findSecondMinimumValue(&x))
}

func findSecondMinimumValue(root *TreeNode) int {
	smallest := math.MaxInt
	second := math.MaxInt

	traverse(root, &smallest, &second)

	if second == math.MaxInt {
		return -1
	}

	return second
}

func traverse(node *TreeNode, smallest *int, second *int) {
	if node.Val != *smallest && node.Val != *second {
		if node.Val < *smallest {
			*second = *smallest
			*smallest = node.Val
		} else if node.Val < *second {
			*second = node.Val
		}
	}

	if node.Left != nil {
		traverse(node.Left, smallest, second)
	}

	if node.Right != nil {
		traverse(node.Right, smallest, second)
	}
}
