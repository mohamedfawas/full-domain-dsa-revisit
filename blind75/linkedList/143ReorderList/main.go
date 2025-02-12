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
func reorderList(head *ListNode) {
	// Step 1: Handle base cases
	// If the list is empty, has only one node, or has only two nodes,
	// no reordering is needed
	if head == nil || head.Next == nil || head.Next.Next == nil {
		return
	}

	// Step 2: Find the middle of the linked list using slow and fast pointers
	// - slow pointer moves one step at a time
	// - fast pointer moves two steps at a time
	// When fast reaches the end, slow will be at the middle
	slow, fast := head, head

	for fast.Next != nil && fast.Next.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	// Example: For list [1,2,3,4,5]
	// After this loop:
	// - slow points to node with value 3 (middle)
	// - fast points to node with value 5 (end)

	// Step 3: Split the list into two halves
	secondHalf := slow.Next // Points to the first node of second half
	slow.Next = nil         // Terminates the first half
	// Example:
	// First half: 1->2->3->nil
	// Second half: 4->5->nil

	// Step 4: Reverse the second half of the list
	var prev *ListNode = nil
	curr := secondHalf

	for curr != nil {
		front := curr.Next // Save the next node
		curr.Next = prev   // Reverse the current pointer
		prev = curr        // Move prev to current node
		curr = front       // Move to the next node
	}
	// Example: Second half after reversal: 5->4->nil
	// prev now points to the head of reversed second half (5)

	// Step 5: Merge the two halves by interleaving nodes
	firstHalf := head // Points to the first node of first half
	secondHalf = prev // Points to the first node of reversed second half

	for secondHalf != nil {
		// Save the next nodes in both lists
		firstNext := firstHalf.Next
		secondNext := secondHalf.Next

		// Connect nodes from first half to second half
		firstHalf.Next = secondHalf

		// Connect nodes from second half to next node in first half
		secondHalf.Next = firstNext

		// Move pointers to next positions for next iteration
		firstHalf = firstNext
		secondHalf = secondNext
	}
	// Example: With [1,2,3] and [5,4]
	// 1. firstHalf(1)->secondHalf(5), then firstHalf=2, secondHalf=4
	// 2. firstHalf(2)->secondHalf(4), then firstHalf=3, secondHalf=nil
	// 3. Loop ends, final list: 1->5->2->4->3->nil
}
