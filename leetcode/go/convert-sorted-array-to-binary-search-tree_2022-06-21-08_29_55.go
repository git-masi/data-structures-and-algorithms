// https://leetcode.com/problems/convert-sorted-array-to-binary-search-tree/

// I failed this attempt but I looked up how to balance a BST, which I had
// never done before, so that was a positive. I'll need to revisit this later.

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
	x := []int{0, 1, 2, 3, 4, 5}

	fmt.Println(sortedArrayToBST(x))
}

func sortedArrayToBST(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}

	mid := len(nums) / 2

	root := TreeNode{
		Val:   nums[mid],
		Left:  nil,
		Right: nil,
	}

	if len(nums) == 1 {
		return &root
	}

	curr := &root

	for i := mid - 1; i >= 0; i-- {
		curr.Left = &TreeNode{
			Val:   nums[i],
			Left:  nil,
			Right: nil,
		}

		curr = curr.Left
	}

	curr = &root

	for i := mid + 1; i < len(nums); i++ {
		curr.Right = &TreeNode{
			Val:   nums[i],
			Left:  nil,
			Right: nil,
		}

		curr = curr.Right
	}

	return &root
}
