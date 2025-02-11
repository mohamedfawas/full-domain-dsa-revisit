package main

type ListNode struct {
	Val  int
	Next *ListNode
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func hasCycle(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return false
	}

	slow := head // tortoise pointer
	fast := head // hare pointer

	for fast != nil && fast.Next != nil {
		// Move slow pointer one step
		slow = slow.Next

		// Move fast pointer two steps
		fast = fast.Next.Next

		// Step 4: Check if pointers meet
		// If they meet, there is a cycle
		if slow == fast {
			return true
		}
	}

	return false
}
