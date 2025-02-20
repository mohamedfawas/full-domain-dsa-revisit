package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func kthSmallest(root *TreeNode, k int) int {
	// This slice will store the values of the nodes in sorted order.
	var values []int

	// Define a helper function to perform in-order traversal.
	// In-order traversal visits left subtree, then root, then right subtree.
	var traverse func(*TreeNode)
	traverse = func(node *TreeNode) {
		if node == nil {
			return
		}

		// Visit left subtree first to get smaller values.
		traverse(node.Left)

		// Append the current node's value to the slice.
		// Example: If node.Val is 2, append 2 to values.
		values = append(values, node.Val)

		// Visit right subtree to get larger values.
		traverse(node.Right)
	}

	// Start the in-order traversal from the root node.
	traverse(root)

	// The kth smallest element is at position k-1 in the sorted slice.
	// Example: For k=1, return values[0].
	return values[k-1]
}
