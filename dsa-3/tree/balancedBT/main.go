package main

import "fmt"

// Node represents a node in an AVL tree
type Node struct {
	value  int
	left   *Node
	right  *Node
	height int
}

// Utility functions

//Get height of node
func height(n *Node) int {
	if n == nil {
		return 0
	}
	return n.height
}

// Get Balance Factor
func getBalance(n *Node) int {
	if n == nil {
		return 0
	}
	return height(n.left) - height(n.right)
}

// Create a New Node
func newNode(value int) *Node {
	return &Node{value: value, height: 1}
}

// Step 3: Rotations to Maintain Balance
// rightRotate performs a right rotation on an AVL tree node.
// Used in the Left-Left (LL) imbalance case when a node has a balance factor > 1
// and the left subtree is left-heavy.
//
// Example:
//
//       y                   x
//      / \                 / \
//     x   T3     -->      T1  y
//    / \                     / \
//   T1  T2                  T2  T3
//
// Where T1, T2, T3 are subtrees.
//
// Before rotation:
// - Node y is the root with balance factor > 1 (left-heavy)
// - Node x is y's left child
// - T1 is x's left subtree
// - T2 is x's right subtree
// - T3 is y's right subtree
//
// After rotation:
// - Node x becomes the new root
// - T1 remains x's left subtree
// - Node y becomes x's right child
// - T2 becomes y's left subtree
// - T3 remains y's right subtree
func rightRotate(y *Node) *Node {
	// Step 1: Store references to the nodes that will be moving
	x := y.left   // x is the left child of y that will become the new root
	T2 := x.right // T2 is the right subtree of x that will need to be reattached

	// Step 2: Perform the rotation
	// Make y the right child of x (changes the parent-child relationship)
	x.right = y
	// Attach T2 as the left child of y (maintain BST property: T2 values between x and y)
	y.left = T2

	// Step 3: Update the height values
	// Height of a node is 1 + the maximum height of its children
	// We update y first because x's height now depends on y's updated height
	y.height = max(height(y.left), height(y.right)) + 1
	x.height = max(height(x.left), height(x.right)) + 1

	// Step 4: Return the new root of this subtree
	return x

	/*
		Numerical Example:

		Before rotation:
		    30 (y, height=3)
		    /\
		   20 (x, height=2)  40
		  /\
		 10  25 (T2)

		After rotation:
		      20 (x, height=2)
		     /  \
		    10   30 (y, height=2)
		        /  \
		       25   40
	*/
}

// leftRotate performs a left rotation on an AVL tree node.
// Used in the Right-Right (RR) imbalance case when a node has a balance factor < -1
// and the right subtree is right-heavy.
//
// Example:
//
//     x                           y
//    / \                         / \
//   T1  y         -->           x   T3
//      / \                     / \
//     T2  T3                  T1  T2
//
// Where T1, T2, T3 are subtrees.
//
// Before rotation:
// - Node x is the root with balance factor < -1 (right-heavy)
// - Node y is x's right child
// - T1 is x's left subtree
// - T2 is y's left subtree
// - T3 is y's right subtree
//
// After rotation:
// - Node y becomes the new root
// - Node x becomes y's left child
// - T1 remains x's left subtree
// - T2 becomes x's right subtree
// - T3 remains y's right subtree
func leftRotate(x *Node) *Node {
	// Step 1: Store references to the nodes that will be moving
	y := x.right // y is the right child of x that will become the new root
	T2 := y.left // T2 is the left subtree of y that will need to be reattached

	// Step 2: Perform the rotation
	// Make x the left child of y (changes the parent-child relationship)
	y.left = x
	// Attach T2 as the right child of x (maintain BST property: T2 values between x and y)
	x.right = T2

	// Step 3: Update the height values
	// Height of a node is 1 + the maximum height of its children
	// We update x first because y's height now depends on x's updated height
	x.height = max(height(x.left), height(x.right)) + 1
	y.height = max(height(y.left), height(y.right)) + 1

	// Step 4: Return the new root of this subtree
	return y

	/*
		Numerical Example:

		Before rotation:
		    20 (x, height=3)
		   /  \
		  10   30 (y, height=2)
		      /  \
		     25   40

		After rotation:
		        30 (y, height=2)
		       /  \
		      20   40
		     /  \
		    10  25
	*/
}

//  Step 4: Insert a Node (with Balancing)
// insert adds a new value to the AVL tree while maintaining balance.
// This function recursively traverses the tree, inserts the new value in the correct position,
// updates heights, and performs rotations as needed to keep the tree balanced.
//
// The function handles four possible imbalance cases:
// 1. Left-Left (LL): Node is left-heavy and value was inserted in left subtree's left branch
// 2. Right-Right (RR): Node is right-heavy and value was inserted in right subtree's right branch
// 3. Left-Right (LR): Node is left-heavy and value was inserted in left subtree's right branch
// 4. Right-Left (RL): Node is right-heavy and value was inserted in right subtree's left branch
//
// Parameters:
// - node: Current node being processed in the recursion
// - value: Value to be inserted
//
// Returns:
// - Updated node reference after insertion and potential rebalancing
func insert(node *Node, value int) *Node {
	// Step 1: Perform standard BST insertion
	// Base case: If we've reached a nil node, create a new node with the value
	if node == nil {
		return newNode(value) // Create and return a new leaf node with height 1
	}

	// Recursive insertion based on BST property (smaller values go left, larger go right)
	if value < node.value {
		// If value is less than current node's value, insert in left subtree
		node.left = insert(node.left, value)
	} else if value > node.value {
		// If value is greater than current node's value, insert in right subtree
		node.right = insert(node.right, value)
	} else {
		// If value already exists, return the current node unchanged (no duplicates)
		return node
	}

	// Step 2: Update height of current node
	// Height is 1 plus the maximum height of the left and right subtrees
	node.height = max(height(node.left), height(node.right)) + 1

	// Step 3: Calculate balance factor to check if node became unbalanced
	// Balance factor = height of left subtree - height of right subtree
	balance := getBalance(node)

	// Step 4: Perform necessary rotations if the node is unbalanced

	// Case 1: Left-Left (LL) imbalance
	// The node is left-heavy (balance > 1) and the new value was inserted in the left subtree's left branch
	//
	// Example:
	//      z                     y
	//     / \                   / \
	//    y   T4   rightRotate  x   z
	//   / \      ----------->  /\  /\
	//  x   T3                 T1 T2 T3 T4
	// / \
	//T1  T2
	if balance > 1 && value < node.left.value {
		return rightRotate(node) // Single right rotation
	}

	// Case 2: Right-Right (RR) imbalance
	// The node is right-heavy (balance < -1) and the new value was inserted in the right subtree's right branch
	//
	// Example:
	//    z                          y
	//   / \                        / \
	//  T1  y       leftRotate     z   x
	//     / \     ----------->   / \ / \
	//    T2  x                  T1 T2 T3 T4
	//       / \
	//      T3 T4
	if balance < -1 && value > node.right.value {
		return leftRotate(node) // Single left rotation
	}

	// Case 3: Left-Right (LR) imbalance
	// The node is left-heavy (balance > 1) and the new value was inserted in the left subtree's right branch
	//
	// Example:
	//      z                        z                       x
	//     / \                      / \                    /   \
	//    y   T4    leftRotate     x   T4   rightRotate  y      z
	//   / \        ----------->  / \      ----------->  / \    / \
	//  T1  x                    y   T3                 T1 T2  T3 T4
	//     / \                  / \
	//    T2 T3                T1 T2
	if balance > 1 && value > node.left.value {
		// First left rotate the left child to transform LR to LL case
		node.left = leftRotate(node.left)
		// Then right rotate the root node
		return rightRotate(node)
	}

	// Case 4: Right-Left (RL) imbalance
	// The node is right-heavy (balance < -1) and the new value was inserted in the right subtree's left branch
	//
	// Example:
	//      z                          z                          x
	//     / \                        / \                       /   \
	//    T1  y     rightRotate      T1  x       leftRotate   z      y
	//       / \    ----------->        / \     -----------> / \    / \
	//      x  T4                      T2  y                T1 T2  T3 T4
	//     / \                            / \
	//    T2 T3                          T3 T4
	if balance < -1 && value < node.right.value {
		// First right rotate the right child to transform RL to RR case
		node.right = rightRotate(node.right)
		// Then left rotate the root node
		return leftRotate(node)
	}

	// Step 5: Return the unchanged node pointer if no rotations were needed
	return node

	/*
		Numerical Example of insert and rebalance:

		1. Initial tree:
		       30
		      /  \
		    20    40

		2. Insert 10 (LL case):
		       30
		      /  \
		    20    40
		   /
		  10

		  After rebalance:
		       20
		      /  \
		    10    30
		           \
		           40

		3. Insert 25 (RL case):
		       20
		      /  \
		    10    30
		          /  \
		        25    40

		  No rebalance needed (tree is balanced)
	*/
}

// Step 5: In-Order Traversal (For Display)
func inOrder(node *Node) {
	if node != nil {
		inOrder(node.left)
		fmt.Printf("%d ", node.value)
		inOrder(node.right)
	}
}

func main() {
	var root *Node

	// Insert elements
	values := []int{10, 20, 30, 40, 50, 25}
	for _, v := range values {
		root = insert(root, v)
	}

	fmt.Println("In-order traversal of the balanced AVL tree:")
	inOrder(root) // Should print a sorted sequence
}
