package main

// Definition for a binary tree node
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func invertTree(root *TreeNode) *TreeNode {
	// Base case: if the root is nil (empty tree or leaf node's child), return nil
	// This is our terminating condition for recursion
	if root == nil {
		return nil
	}

	// Store the left subtree in a temporary variable
	// We need this because we'll be overwriting root.Left
	// Example: In [4,2,7,1,3,6,9], for root=4, temp will store the entire left subtree starting with 2
	temp := root.Left

	// Step 1: Assign the inverted right subtree to the left child
	// We recursively invert the right subtree before assigning
	// Example: For root=4, root.Left will become the inverted subtree of [7,6,9]
	root.Left = invertTree(root.Right)

	// Step 2: Assign the inverted left subtree (which we stored in temp) to the right child
	// Example: For root=4, root.Right will become the inverted subtree of [2,1,3]
	root.Right = invertTree(temp)

	// Return the root node of the inverted tree
	return root
}

/*
Let's walk through Example 1 step by step:

Initial tree:
     4
   /   \
  2     7
 / \   / \
1   3 6   9

Step 1: Start at root (4)
- Store left subtree (2,1,3) in temp
- Recursively invert right subtree (7,6,9)
- Recursively invert left subtree (2,1,3)

When inverting (7,6,9):
- Store 6 in temp
- Set left to inverted 9 (leaf node)
- Set right to inverted 6 (leaf node)
Result: 7,9,6

When inverting (2,1,3):
- Store 1 in temp
- Set left to inverted 3 (leaf node)
- Set right to inverted 1 (leaf node)
Result: 2,3,1

Final result:
     4
   /   \
  7     2
 / \   / \
9   6 3   1

Time Complexity: O(n) where n is the number of nodes in the tree
Space Complexity: O(h) where h is the height of the tree, due to recursion stack
*/
