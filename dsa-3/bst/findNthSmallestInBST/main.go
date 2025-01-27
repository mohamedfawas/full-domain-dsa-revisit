// Find Kth smallest element in a bst

package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func kthSmallest(root *TreeNode, k int) int {
	// We need to track two values across recursive calls:
	count := 0                                 // 1. How many nodes we've visited so far
	result := 0                                // 2. The final answer we want to return
	inOrderTraversal(root, k, &count, &result) // Pass addresses using &
	return result
}

func inOrderTraversal(node *TreeNode, k int, count *int, result *int) {
	// Base case: stop if we reach a nil node or already found the answer
	if node == nil || *count >= k {
		return
	}

	// 1. FIRST, DRILL DOWN TO THE LEFT-MOST NODE (SMALLEST VALUE)
	inOrderTraversal(node.Left, k, count, result)

	// 2. PROCESS CURRENT NODE ONLY IF WE HAVEN'T FOUND THE ANSWER YET
	if *count < k {
		*count++ // Increment visited node count

		// Check if this is the k-th smallest node
		if *count == k {
			*result = node.Val // Store the answer
			return             // Stop further exploration
		}

		// 3. ONLY EXPLORE RIGHT SUBTREE IF STILL NEEDING MORE NODES
		inOrderTraversal(node.Right, k, count, result)
	}
}
