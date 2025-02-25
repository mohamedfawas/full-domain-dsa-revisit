package main

import "fmt"

// Node represents a node in the ternary tree
// Each node contains a value and pointers to left, middle, and right children
type Node struct {
	Value  int
	Left   *Node
	Middle *Node
	Right  *Node
}

// NewNode creates and returns a new node with the given value
func NewNode(value int) *Node {
	return &Node{
		Value:  value,
		Left:   nil,
		Middle: nil,
		Right:  nil,
	}
}

// TernaryTree represents the complete tree structure
type TernaryTree struct {
	Root *Node
}

// NewTernaryTree creates and returns a new empty ternary tree
func NewTernaryTree() *TernaryTree {
	return &TernaryTree{Root: nil}
}

// Insert adds a new value to the tree using level-order traversal
// This ensures the tree is filled from left to right, level by level
func (t *TernaryTree) Insert(value int) {
	newNode := NewNode(value)

	// If tree is empty, make the new node the root
	if t.Root == nil {
		t.Root = newNode
		return
	}

	// Use a queue to perform level-order traversal
	queue := []*Node{t.Root}

	// Continue until we find a node with an empty child position
	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:] // Remove the first element

		// Try to insert in left position
		if current.Left == nil {
			current.Left = newNode
			return
		}
		queue = append(queue, current.Left)

		// Try to insert in middle position
		if current.Middle == nil {
			current.Middle = newNode
			return
		}
		queue = append(queue, current.Middle)

		// Try to insert in right position
		if current.Right == nil {
			current.Right = newNode
			return
		}
		queue = append(queue, current.Right)
	}
}

// PrintLevelOrder prints the tree level by level
func (t *TernaryTree) PrintLevelOrder() {
	if t.Root == nil {
		fmt.Println("Empty tree")
		return
	}

	// Use a queue for level-order traversal
	queue := []*Node{t.Root}
	level := 0

	for len(queue) > 0 {
		// Print current level
		fmt.Printf("Level %d: ", level)
		levelSize := len(queue)

		for i := 0; i < levelSize; i++ {
			current := queue[0]
			queue = queue[1:] // Remove the first element
			fmt.Printf("%d ", current.Value)

			// Add children to queue if they exist
			if current.Left != nil {
				queue = append(queue, current.Left)
			}
			if current.Middle != nil {
				queue = append(queue, current.Middle)
			}
			if current.Right != nil {
				queue = append(queue, current.Right)
			}
		}
		fmt.Println() // New line after each level
		level++
	}
}

func main() {
	// Example usage of the ternary tree
	tree := NewTernaryTree()

	// Insert some values
	values := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	for _, v := range values {
		tree.Insert(v)
	}

	// Print the tree level by level
	fmt.Println("Ternary Tree Level Order Traversal:")
	tree.PrintLevelOrder()

	/* Output will be:
	   Ternary Tree Level Order Traversal:
	   Level 0: 1
	   Level 1: 2 3 4
	   Level 2: 5 6 7 8 9 10
	*/
}
