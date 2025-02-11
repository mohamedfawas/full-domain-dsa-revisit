package main

// TreeNode represents a node in a binary tree
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSameTree(p *TreeNode, q *TreeNode) bool {
	// Base Case 1: If both nodes are nil, we've reached the end of both trees
	// at the same time, which means this path is identical
	// Example: comparing two empty subtrees
	//   p: nil, q: nil -> return true
	if p == nil && q == nil {
		return true
	}

	// Base Case 2: If one node is nil and the other isn't, trees are different
	// Example:
	//   p: [1,2], q: [1,null,2]
	//   When comparing p.Left (2) with q.Left (nil) -> return false
	if (p == nil && q != nil) || (p != nil && q == nil) {
		return false
	}

	// Base Case 3: If the values of current nodes are different, trees are different
	// Example:
	//   p: [1,2,1], q: [1,1,2]
	//   When comparing p.Left.Val (2) with q.Left.Val (1) -> return false
	if p.Val != q.Val {
		return false
	}

	// Recursive Case: If we've made it here, current nodes are the same
	// Now we need to check both left and right subtrees
	// We use && because BOTH subtrees must be identical for the trees to be same
	//
	// Example of recursive flow for p = [1,2,3], q = [1,2,3]:
	// 1. Check root nodes (1 == 1)
	// 2. Recursively check left subtrees (2 == 2)
	// 3. Recursively check right subtrees (3 == 3)
	// 4. All checks pass -> return true
	return isSameTree(p.Left, q.Left) && isSameTree(p.Right, q.Right)
}

// Example usage:
// tree1 := &TreeNode{
//     Val: 1,
//     Left: &TreeNode{Val: 2},
//     Right: &TreeNode{Val: 3},
// }
//
// tree2 := &TreeNode{
//     Val: 1,
//     Left: &TreeNode{Val: 2},
//     Right: &TreeNode{Val: 3},
// }
//
// result := isSameTree(tree1, tree2) // returns true
