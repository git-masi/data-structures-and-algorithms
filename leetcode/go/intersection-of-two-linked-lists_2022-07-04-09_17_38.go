// https://leetcode.com/problems/intersection-of-two-linked-lists/

// This was pretty easy but I'm not sure if the solution is optimal
// I know that this will run in O(n + m) time and requires O(n) space
// but is there a way to do this without extra space and the same time
// time complexity?

package main

import (
	"fmt"
)

func main() {
	common := ListNode{
		Val: 8,
		Next: &ListNode{
			Val: 4,
			Next: &ListNode{
				Val: 1,
			},
		},
	}

	a := ListNode{
		Val:  2,
		Next: &common,
	}

	b := ListNode{
		Val: 7,
		Next: &ListNode{
			Val: 1,
			Next: &ListNode{
				Val:  9,
				Next: &common,
			},
		},
	}

	fmt.Println(getIntersectionNode(&a, &b))
	fmt.Println(getIntersectionNode(&a, &b) == &common)
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	if headA == nil || headB == nil {
		return nil
	}

	m := map[*ListNode]bool{}
	var result *ListNode

	for {
		m[headA] = true

		if headA.Next == nil {
			break
		}

		headA = headA.Next
	}

	for {
		_, ok := m[headB]

		if ok {
			result = headB
			break
		}

		if headB.Next == nil {
			break
		}

		headB = headB.Next
	}

	return result
}
