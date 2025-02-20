package main

// TreeNode represents a node in a binary tree
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func buildTree(preorder []int, inorder []int) *TreeNode {
	// Create a map to store the indices of values in the inorder array for quick lookup.
	indexMap := make(map[int]int)
	for i, val := range inorder {
		indexMap[val] = i
	}

	// Initialize a pointer to track the current position in the preorder array.
	preIndex := 0

	// Define a helper function to recursively build the tree using divide and conquer.
	var helper func(inStart, inEnd int) *TreeNode
	helper = func(inStart, inEnd int) *TreeNode {
		// Base case: if there are no elements to construct the subtree, return nil.
		if inStart > inEnd {
			return nil
		}

		// The current root value is the next element in the preorder array.
		rootVal := preorder[preIndex]
		root := &TreeNode{Val: rootVal}
		preIndex++ // Move to the next element for subsequent calls.

		// Find the position of the current root in the inorder array.
		inIndex := indexMap[rootVal]

		// Recursively build the left subtree using elements left of the root in inorder.
		root.Left = helper(inStart, inIndex-1)
		// Recursively build the right subtree using elements right of the root in inorder.
		root.Right = helper(inIndex+1, inEnd)

		return root
	}

	// Start the recursion with the full range of the inorder array.
	return helper(0, len(inorder)-1)
}
