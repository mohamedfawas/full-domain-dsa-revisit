package main

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSubtree(root *TreeNode, subRoot *TreeNode) bool {
	// If the main tree (root) is nil, there can't be any subtree
	if root == nil {
		return false
	}

	// First, check if the current root node could be the start of our subtree
	// by calling isSameTree to compare the entire structure from this point
	if isSameTree(root, subRoot) {
		return true
	}

	// If the current node isn't the start of our subtree, recursively check:
	// 1. If the subtree exists in the left branch OR
	// 2. If the subtree exists in the right branch
	return isSubtree(root.Left, subRoot) || isSubtree(root.Right, subRoot)
}

// isSameTree checks if two trees are exactly identical in structure and values
func isSameTree(p *TreeNode, q *TreeNode) bool {
	// Case 1: Both nodes are nil
	// Example: When comparing two empty subtrees
	// p: nil, q: nil -> return true
	if p == nil && q == nil {
		return true
	}

	// Case 2: One node is nil and other isn't
	// Example: p: [1,2,3], q: [1,2] -> return false
	// Because one tree has a node where the other has nil
	if p == nil || q == nil {
		return false
	}

	// Case 3: Values don't match
	// Example: p: [1,2,3], q: [1,2,4] -> return false
	// Because 3 != 4
	if p.Val != q.Val {
		return false
	}

	// Case 4: Current nodes match, so recursively check both left and right subtrees
	// Example: p: [1,2,3], q: [1,2,3]
	// 1. First checks if 1==1 (root values match)
	// 2. Then recursively checks if left subtrees match (2==2)
	// 3. And if right subtrees match (3==3)
	// Only returns true if ALL checks pass
	return isSameTree(p.Left, q.Left) && isSameTree(p.Right, q.Right)
}
