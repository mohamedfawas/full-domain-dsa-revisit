package main

import (
	"fmt"
	"math"
)

// First, we'll create a structure for our tree node
// Think of it like a box that can hold a number and point to two other boxes
type TreeNode struct {
	Value int       // The number inside our box
	Left  *TreeNode // Points to the left box (should be smaller numbers)
	Right *TreeNode // Points to the right box (should be bigger numbers)
}

// This function will help us check if our tree follows BST rules
// It's like having a helper that checks if numbers are in the correct order
func isBST(root *TreeNode) bool {
	// We'll use a helper function that also checks if numbers are within allowed range
	return isBSTHelper(root, math.MinInt64, math.MaxInt64)
}

// This is our helper function that does the actual checking
// min: smallest allowed number for current node
// max: biggest allowed number for current node
func isBSTHelper(node *TreeNode, min, max int) bool {
	// If we reach an empty box (nil node), it's okay!
	// Empty boxes don't break any rules
	if node == nil {
		return true
	}

	// Rule 1: Check if current number is within allowed range
	// It's like checking if a number in a sequence is in the right place
	if node.Value <= min || node.Value >= max {
		return false
	}

	// Rule 2: Check left side
	// All numbers on left should be smaller than current number
	// So we update the max limit to current number
	if !isBSTHelper(node.Left, min, node.Value) {
		return false
	}

	// Rule 3: Check right side
	// All numbers on right should be bigger than current number
	// So we update the min limit to current number
	if !isBSTHelper(node.Right, node.Value, max) {
		return false
	}

	// If we reach here, all rules are followed!
	return true
}

// Here's how to use our functions
func main() {
	// Let's create a valid BST first
	//       5
	//      / \
	//     3   7
	//    / \
	//   2   4

	validBST := &TreeNode{
		Value: 5,
		Left: &TreeNode{
			Value: 3,
			Left:  &TreeNode{Value: 2},
			Right: &TreeNode{Value: 4},
		},
		Right: &TreeNode{
			Value: 7,
		},
	}

	// Let's create an invalid BST
	//       5
	//      / \
	//     3   7
	//    / \
	//   2   6  <- This 6 makes it invalid because it's bigger than 5

	invalidBST := &TreeNode{
		Value: 5,
		Left: &TreeNode{
			Value: 3,
			Left:  &TreeNode{Value: 2},
			Right: &TreeNode{Value: 6}, // This breaks BST property
		},
		Right: &TreeNode{
			Value: 7,
		},
	}

	// Check both trees
	fmt.Println("Is valid BST correct?:", isBST(validBST))     // Should print: true
	fmt.Println("Is invalid BST correct?:", isBST(invalidBST)) // Should print: false
}
