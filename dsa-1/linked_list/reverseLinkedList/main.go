// Reverse the given linked list

package main

import "fmt"

type Node struct {
	Value int
	Next  *Node
}

func reverseLinkedList(head *Node) *Node {
	var prev *Node
	current := head
	for current != nil {
		next := current.Next
		current.Next = prev
		prev = current
		current = next
	}

	return prev
}

func main() {
	// Create a linked list
	head := &Node{Value: 1}
	head.Next = &Node{Value: 2}
	head.Next.Next = &Node{Value: 3}
	head.Next.Next.Next = &Node{Value: 4}
	head.Next.Next.Next.Next = &Node{Value: 5}

	// Reverse the linked list
	head = reverseLinkedList(head)

	fmt.Println("Reversed linked list:")
	// Print the reversed linked list
	for head != nil {
		println(head.Value)
		head = head.Next
	}
}
