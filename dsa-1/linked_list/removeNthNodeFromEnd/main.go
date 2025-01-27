// Remove nth node from the end of the linkedlist

package main

import "fmt"

type Node struct {
	Value int
	Next  *Node
}

func removeNthNode(head *Node, n int) *Node {
	// Create a dummy node to handle the edge case (like removing head node)

	dummy := &Node{Next: head}

	fast, slow := dummy, dummy

	// Move fast pointer n steps ahead
	for i := 0; i < n; i++ {
		fast = fast.Next
	}

	// Move both pointers until fast reaches the last node
	for fast.Next != nil {
		fast = fast.Next
		slow = slow.Next
	}

	// Remove the nth node by skipping it
	slow.Next = slow.Next.Next

	return dummy.Next
}

func main() {
	// Create a linkedlist
	head := &Node{Value: 1}
	head.Next = &Node{Value: 2}
	head.Next.Next = &Node{Value: 3}
	head.Next.Next.Next = &Node{Value: 4}
	head.Next.Next.Next.Next = &Node{Value: 5}
	fmt.Println("Created the linkedlist: 1->2->3->4->5")

	// Remove 2nd node from the end
	head = removeNthNode(head, 2)

	fmt.Println("After removing 2nd node from the end: ")
	// Print the linkedlist
	for head != nil {
		println(head.Value)
		head = head.Next
	}
}
