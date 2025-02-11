package main

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxDepth(root *TreeNode) int {
	// Base case: if the root is nil (empty tree), return 0
	// This handles both empty trees and when we reach beyond leaf nodes
	if root == nil {
		return 0
	}

	// Recursively calculate the depth of left and right subtrees
	leftDepth := maxDepth(root.Left)   // Get max depth of left subtree
	rightDepth := maxDepth(root.Right) // Get max depth of right subtree

	// Return the larger depth between left and right subtrees, plus 1 for the current node
	// We use max() to compare the depths and add 1 to include the current level
	return max(leftDepth, rightDepth) + 1
}

// Helper function to find maximum of two integers
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
