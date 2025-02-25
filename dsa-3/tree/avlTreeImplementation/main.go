package main

import "fmt"

// Node represents a node in the AVL tree
type Node struct {
	Value     int   // Value stored in the node
	Height    int   // Height of the node for balancing
	LeftNode  *Node // Left child pointer
	RightNode *Node // Right child pointer
}

// AVLTree represents the tree structure
type AVLTree struct {
	Root *Node // Root node of the tree
}

// getHeight returns the height of a node
// Returns 0 for nil nodes
func getHeight(node *Node) int {
	if node == nil {
		return 0
	}
	return node.Height
}

// getBalance returns the balance factor of a node
// Balance factor = height of left subtree - height of right subtree
func getBalance(node *Node) int {
	if node == nil {
		return 0
	}
	return getHeight(node.LeftNode) - getHeight(node.RightNode)
}

// updateHeight updates the height of a node based on its children
func (node *Node) updateHeight() {
	leftHeight := getHeight(node.LeftNode)
	rightHeight := getHeight(node.RightNode)

	// Height is 1 + the maximum height of left or right subtree
	if leftHeight > rightHeight {
		node.Height = leftHeight + 1
	} else {
		node.Height = rightHeight + 1
	}
}

// rotateRight performs a right rotation on the given node
// Example: Consider this tree structure before rotation:
//       y
//      / \
//     x   c
//    / \
//   a   T2
//
// After right rotation it becomes:
//       x
//      / \
//     a   y
//        / \
//       T2  c
//
// Where:
// - y is the node we're rotating around
// - x is y's left child
// - T2 is x's right child (which will move)
// - a and c are subtrees that don't move
func rotateRight(y *Node) *Node {
	// Step 1: Save the nodes we'll be moving around
	// x will become the new root of this subtree
	x := y.LeftNode
	// T2 is the right child of x that will need to move
	T2 := x.RightNode

	// Step 2: Perform the rotation
	// Make y the right child of x
	x.RightNode = y
	// Move T2 to be the left child of y
	// (T2 needs to go here because it's greater than x but less than y)
	y.LeftNode = T2

	// Step 3: Update the heights of the modified nodes
	// We update y first because it's now lower in the tree
	y.updateHeight()
	// Then update x because it depends on y's new height
	x.updateHeight()

	// Step 4: Return the new root of this subtree (x)
	return x
}

// Example case where right rotation is needed:
// Initial tree after inserting 30, 20, 10:
//      30   (balance = 2)
//     /
//    20     (balance = 1)
//   /
//  10       (balance = 0)
//
// After right rotation at node 30:
//      20   (balance = 0)
//     /  \
//    10   30 (balance = 0)

// rotateLeft performs a left rotation on the given node
func rotateLeft(x *Node) *Node {
	y := x.RightNode
	T2 := y.LeftNode

	// Perform rotation
	y.LeftNode = x
	x.RightNode = T2

	// Update heights
	x.updateHeight()
	y.updateHeight()

	return y
}

// Insert adds a new value to the AVL tree
func (tree *AVLTree) Insert(value int) {
	tree.Root = insert(tree.Root, value)
}

// insert recursively inserts a value into the tree and balances it
func insert(node *Node, value int) *Node {
	// Step 1: Perform normal BST insertion
	if node == nil {
		return &Node{Value: value, Height: 1}
	}

	if value < node.Value {
		node.LeftNode = insert(node.LeftNode, value)
	} else if value > node.Value {
		node.RightNode = insert(node.RightNode, value)
	} else {
		// Duplicate values are not allowed
		return node
	}

	// Step 2: Update height of current node
	node.updateHeight()

	// Step 3: Get the balance factor
	balance := getBalance(node)

	// Step 4: Balance the tree if needed

	// Case 1: Left Left Case (LL Rotation)
	// When: balance > 1 (left subtree is higher) AND new value is less than left child
	// Example: Inserting 10 into this tree:
	//      30 (balance=2)          20
	//     /                       /  \
	//    20 (balance=1)    =>   10   30
	//   /
	//  10
	// This is fixed with a single right rotation at node 30
	if balance > 1 && value < node.LeftNode.Value {
		return rotateRight(node)
	}

	// Case 2: Right Right Case (RR Rotation)
	// When: balance < -1 (right subtree is higher) AND new value is greater than right child
	// Example: Inserting 30 into this tree:
	//    10 (balance=-2)          20
	//      \                     /  \
	//       20 (balance=-1) => 10   30
	//        \
	//         30
	// This is fixed with a single left rotation at node 10
	if balance < -1 && value > node.RightNode.Value {
		return rotateLeft(node)
	}

	// Case 3: Left Right Case (LR Rotation)
	// When: balance > 1 (left subtree is higher) AND new value is greater than left child
	// Example: Inserting 20 into this tree:
	//      30 (balance=2)     30 (balance=2)       20
	//     /                  /                     /  \
	//    10 (balance=-1) => 20          =>       10   30
	//      \              /
	//       20          10
	// This requires two steps:
	// 1. Left rotation at node 10 (left child)
	// 2. Right rotation at node 30 (root)
	if balance > 1 && value > node.LeftNode.Value {
		// First rotate the left subtree to the left
		node.LeftNode = rotateLeft(node.LeftNode)
		// Then rotate the whole tree to the right
		return rotateRight(node)
	}

	// Case 4: Right Left Case (RL Rotation)
	// When: balance < -1 (right subtree is higher) AND new value is less than right child
	// Example: Inserting 20 into this tree:
	//    10 (balance=-2)    10 (balance=-2)        20
	//      \                  \                    /  \
	//       30 (balance=1) =>  20         =>     10   30
	//      /                     \
	//     20                      30
	// This requires two steps:
	// 1. Right rotation at node 30 (right child)
	// 2. Left rotation at node 10 (root)
	if balance < -1 && value < node.RightNode.Value {
		// First rotate the right subtree to the right
		node.RightNode = rotateRight(node.RightNode)
		// Then rotate the whole tree to the left
		return rotateLeft(node)
	}

	return node
}

// InorderTraversal prints the tree values in-order
func (tree *AVLTree) InorderTraversal() {
	inorder(tree.Root)
	fmt.Println()
}

func inorder(node *Node) {
	if node != nil {
		inorder(node.LeftNode)
		fmt.Printf("%d ", node.Value)
		inorder(node.RightNode)
	}
}

func main() {
	// Create a new AVL tree
	tree := &AVLTree{}

	// Example 1: Simple insertions
	fmt.Println("Example 1: Simple insertions")
	tree.Insert(10)
	tree.Insert(20)
	tree.Insert(30)
	fmt.Print("Inorder traversal after inserting 10, 20, 30: ")
	tree.InorderTraversal()
	// Output: 10 20 30

	// Example 2: Creating an unbalanced situation
	fmt.Println("\nExample 2: More complex insertions")
	tree = &AVLTree{} // Reset tree
	values := []int{10, 20, 30, 40, 50, 25}
	fmt.Print("Inserting values: ")
	for _, v := range values {
		fmt.Printf("%d ", v)
		tree.Insert(v)
	}
	fmt.Print("\nFinal inorder traversal: ")
	tree.InorderTraversal()
	// The tree will automatically balance itself
}
