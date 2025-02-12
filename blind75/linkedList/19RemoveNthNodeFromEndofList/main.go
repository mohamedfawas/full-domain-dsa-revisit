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

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	// Initialize fast pointer at the head of the list
	fast := head

	// Move the fast pointer n nodes ahead
	// This creates a gap of exactly n nodes between fast and slow (which we'll initialize later)
	//
	// Example: For list [1,2,3,4,5] and n=2
	// After this loop, fast will be at 3 (the 3rd node)
	for i := 0; i < n; i++ {
		fast = fast.Next
	}

	// Special case: If fast is nil after moving n steps, it means
	// we need to remove the head node (n equals the length of the list)
	//
	// Example: For list [1,2] and n=2, fast would be nil
	// In this case, we simply return head.Next (which might be nil if there was only one node)
	if fast == nil {
		return head.Next
	}

	// Initialize slow pointer at the head
	slow := head

	// Move both pointers at the same pace until fast reaches the last node
	// When fast is at the last node, slow will be at the node just before
	// the one we want to remove
	//
	// Example continued:
	// 1) fast at 3, slow at 1
	// 2) fast at 4, slow at 2
	// 3) fast at 5, slow at 3
	// Now slow.Next (which is 4) is the node we want to remove
	for fast.Next != nil {
		fast = fast.Next
		slow = slow.Next
	}

	// Remove the nth node from the end by updating the Next pointer
	// slow.Next is the node to remove, so we set slow.Next to slow.Next.Next
	//
	// Example continued:
	// slow is at 3, slow.Next is 4 (to be removed)
	// We set slow.Next to slow.Next.Next, which is 5
	// So the list becomes [1,2,3,5]
	slow.Next = slow.Next.Next

	// Return the head of the modified list
	// Note: We can return the original head here because if the head needed to be removed,
	// we would have already returned in the earlier check
	return head
}
