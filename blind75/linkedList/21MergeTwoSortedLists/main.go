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
func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	// Create a dummy node to simplify the merging process.
	dummy := &ListNode{}
	current := dummy

	// Loop through both lists until one is exhausted.
	for list1 != nil && list2 != nil {
		if list1.Val < list2.Val {
			current.Next = list1
			list1 = list1.Next
		} else {
			current.Next = list2
			list2 = list2.Next
		}
		current = current.Next
	}

	// If any elements remain in list1 or list2, append them to the merged list.
	if list1 != nil {
		current.Next = list1
	} else if list2 != nil {
		current.Next = list2
	}

	// The merged list starts from dummy.Next (skip the dummy node).
	return dummy.Next
}
