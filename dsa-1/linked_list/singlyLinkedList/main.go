package main

import "fmt"

// Node represents a single node in the linked list
type Node struct {
	data int
	next *Node
}


// LinkedList represents a singly linked list
type LinkedList struct {
	head *Node
}

// InsertEnd inserts a new node at the end of the list
func (l *LinkedList) InsertEnd(data int) {
	newNode := &Node{data: data}
	if l.head == nil {
		l.head = newNode
		return
	}
	temp := l.head
	for temp.next != nil {
		temp = temp.next
	}
	temp.next = newNode
}


// InsertBeginning inserts a new node at the beginning of the list
func (l *LinkedList) InsertBeginning(data int) {
	newNode := &Node{data: data, next: l.head}
	l.head = newNode
}


// Delete deletes the first occurrence of the node with given data
func (l *LinkedList) Delete(data int) {
	if l.head == nil {
		fmt.Println("List is empty")
		return
	}

	// If head needs to be removed
	if l.head.data == data {
		l.head = l.head.next
		return
	}

	// Search for the node to delete
	temp := l.head
	for temp.next != nil && temp.next.data != data {
		temp = temp.next
	}

	// If found, remove it
	if temp.next != nil {
		temp.next = temp.next.next
	} else {
		fmt.Println("Node not found")
	}
}
// Display prints all elements in the linked list
func (l *LinkedList) Display() {
	if l.head == nil {
		fmt.Println("List is empty")
		return
	}
	temp := l.head
	for temp != nil {
		fmt.Printf("%d -> ", temp.data)
		temp = temp.next
	}
	fmt.Println("nil")
}

func main() {
	list := LinkedList{}

	list.InsertEnd(10)
	list.InsertEnd(20)
	list.InsertEnd(30)
	list.Display() // Output: 10 -> 20 -> 30 -> nil

	list.InsertBeginning(5)
	list.Display() // Output: 5 -> 10 -> 20 -> 30 -> nil

	list.Delete(20)
	list.Display() // Output: 5 -> 10 -> 30 -> nil

	list.Delete(5)
	list.Display() // Output: 10 -> 30 -> nil

	list.Delete(50) // Output: Node not found
}
