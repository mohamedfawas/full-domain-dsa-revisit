package main

import "fmt"

// Node represents a node in the complete binary tree
type Node struct {
	Value int
	Left  *Node
	Right *Node
}

// CompleteBinaryTree represents the tree structure
type CompleteBinaryTree struct {
	Root *Node
}

// NewCompleteBinaryTree creates a new empty complete binary tree
func NewCompleteBinaryTree() *CompleteBinaryTree {
	return &CompleteBinaryTree{
		Root: nil,
	}
}

// Insert adds a new value to the complete binary tree
// It follows level-order insertion to maintain the complete binary tree property
func (t *CompleteBinaryTree) Insert(value int) {
	newNode := &Node{Value: value}

	// If tree is empty, make the new node the root
	if t.Root == nil {
		t.Root = newNode
		return
	}

	// Use a queue to perform level order traversal
	queue := []*Node{t.Root}

	// Continue until we find a node with an empty child position
	for len(queue) > 0 {
		current := queue[0] // Get the first node from queue
		queue = queue[1:]   // Remove the first node from queue

		// If left child is empty, insert here
		if current.Left == nil {
			current.Left = newNode
			return
		}
		// If right child is empty, insert here
		if current.Right == nil {
			current.Right = newNode
			return
		}

		// Add children to queue for next level traversal
		queue = append(queue, current.Left)
		queue = append(queue, current.Right)
	}
}

// PrintLevelOrder prints the tree level by level
func (t *CompleteBinaryTree) PrintLevelOrder() {
	if t.Root == nil {
		fmt.Println("Tree is empty")
		return
	}

	// Use queue for level order traversal
	queue := []*Node{t.Root}

	for len(queue) > 0 {
		levelSize := len(queue)
		for i := 0; i < levelSize; i++ {
			current := queue[0]
			queue = queue[1:]
			fmt.Printf("%d ", current.Value)

			if current.Left != nil {
				queue = append(queue, current.Left)
			}
			if current.Right != nil {
				queue = append(queue, current.Right)
			}
		}
		fmt.Println() // New line for each level
	}
}

func main() {
	// Create a new complete binary tree
	tree := NewCompleteBinaryTree()

	// Insert some values
	values := []int{1, 2, 3, 4, 5, 6, 7}
	fmt.Println("Inserting values:", values)

	for _, val := range values {
		tree.Insert(val)
	}

	fmt.Println("\nPrinting tree level by level:")
	tree.PrintLevelOrder()

	// This will create a tree that looks like:
	//       1
	//     /   \
	//    2     3
	//   / \   / \
	//  4   5 6   7
}
