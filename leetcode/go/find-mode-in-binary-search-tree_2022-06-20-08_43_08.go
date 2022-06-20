// https://leetcode.com/problems/find-mode-in-binary-search-tree/

// I did not solve this HOWEVER I think the problem was poorly written
// basically the problem states that we're dealing with a BST but we're
// not in reality. I even voted down this question and sent the note
// seen below to leetcode.

// My one takeaway from this is to always read the question carefully
// and account for the potential misuse of terms.

// My note to leetcode:
//
// The problem with this question is that the input is not really a BST as most
// people would understand it. Most people follow the definition such that:
// "the key of each internal node being greater than all the keys in the respective
// node's left subtree and less than the ones in its right subtree".
// This question breaks that definition because nodes of the subtrees on either side
// could be equal to the root node. That's not a BST. It's just a binary tree where
// nodes on the right of a given subtree can't be less than the root node.

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
	x := TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 1,
			Left: &TreeNode{
				Val:   1,
				Left:  nil,
				Right: nil,
			},
			Right: nil,
		},
		Right: &TreeNode{
			Val:  2,
			Left: nil,
			Right: &TreeNode{
				Val: 4,
				Left: &TreeNode{
					Val: 4,
					Left: &TreeNode{
						Val: 4,
						Left: &TreeNode{
							Val:   3,
							Left:  nil,
							Right: nil,
						},
						Right: nil,
					},
					Right: nil,
				},
				Right: nil,
			},
		},
	}

	fmt.Println(findMode(&x))
}

func findMode(root *TreeNode) []int {
	mode := []int{}
	maxFreq := 0

	if root == nil {
		return mode
	}

	dfs(root, &maxFreq, 1, &mode)

	return mode
}

func dfs(curr *TreeNode, maxFreq *int, currValCount int, mode *[]int) {
	if curr.Left != nil && curr.Val == curr.Left.Val {
		dfs(curr.Left, maxFreq, currValCount+1, mode)
	} else {
		if currValCount > *maxFreq {
			*maxFreq = currValCount
			*mode = []int{curr.Val}
		} else if currValCount == *maxFreq {
			*mode = append(*mode, curr.Val)
		}

		if curr.Left != nil {
			dfs(curr.Left, maxFreq, 1, mode)
		}
	}

	if curr.Right != nil {
		dfs(curr.Right, maxFreq, 1, mode)
	}
}
