package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val   int
 *     Left  *TreeNode
 *     Right *TreeNode
 * }
 */

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	// Base case: if root is nil or root is either p or q, return root
	// This handles both when we've found one of our target nodes
	// and when we've reached the end of a path without finding anything
	if root == nil || root == p || root == q {
		return root
	}

	/* In a BST, we can use the property that:
	 * - All values in left subtree are less than the current node
	 * - All values in right subtree are greater than the current node
	 *
	 * Example: For tree [6,2,8,0,4,7,9]
	 *       6
	 *      / \
	 *     2   8
	 *    / \ / \
	 *   0  4 7  9
	 */

	// If both p and q values are less than root's value,
	// then LCA must be in the left subtree
	// Example: if p=0 and q=4, and we're at root=6
	// Both 0 and 4 are less than 6, so we go left to 2
	if p.Val < root.Val && q.Val < root.Val {
		return lowestCommonAncestor(root.Left, p, q)
	}

	// If both p and q values are greater than root's value,
	// then LCA must be in the right subtree
	// Example: if p=7 and q=9, and we're at root=6
	// Both 7 and 9 are greater than 6, so we go right to 8
	if p.Val > root.Val && q.Val > root.Val {
		return lowestCommonAncestor(root.Right, p, q)
	}

	/* If we reach here, it means either:
	 * 1. One value is less than root and other is greater
	 *    Example: p=2, q=8, root=6
	 *    This means current root is the LCA
	 * 2. One of the values equals the root value
	 *    Example: p=2, q=4, root=2
	 *    In this case, the root is the LCA
	 */
	return root
}

/* Time Complexity: O(h) where h is the height of the tree
 * - In worst case (skewed tree), h = n where n is number of nodes
 * - In balanced BST, h = log(n)
 *
 * Space Complexity: O(h) for recursive call stack
 *
 * Example walkthrough:
 * Tree:     6
 *          / \
 *         2   8
 *        / \
 *       0   4
 *
 * Case 1: p=2, q=8
 * 1. Start at 6
 * 2. 2 < 6 but 8 > 6, so 6 is LCA
 *
 * Case 2: p=0, q=4
 * 1. Start at 6
 * 2. Both 0 and 4 < 6, go left to 2
 * 3. 0 < 2 but 4 > 2, so 2 is LCA
 */
