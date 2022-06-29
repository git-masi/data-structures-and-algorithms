// https://leetcode.com/problems/linked-list-cycle/

// Overall this is pretty easy using the fast and slow pointer pattern.
// Making sure to account for edge cases is important as always.

package main

import (
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	three := ListNode{
		Val: 3,
	}

	two := ListNode{
		Val: 2,
	}

	three.Next = &two

	two.Next = &ListNode{
		Val: 0,
		Next: &ListNode{
			Val:  -4,
			Next: &two,
		},
	}

	fmt.Println(hasCycle(&three))
}

func hasCycle(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return false
	}

	slow := head
	fast := head.Next

	for {
		if fast == nil || fast.Next == nil {
			break
		}

		if slow == fast {
			return true
		}

		fast = fast.Next.Next
		slow = slow.Next
	}

	return false
}
