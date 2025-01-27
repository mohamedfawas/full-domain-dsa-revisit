package main

// TreeNode represents a node in a Binary Search Tree (BST)
type TreeNode struct {
	Val   int       // Value of the node
	Left  *TreeNode // Pointer to the left child
	Right *TreeNode // Pointer to the right child
}

// kthLargest is the main function to find the kth largest element in the BST
func kthLargest(root *TreeNode, k int) int {
	count := 0  // Counter to keep track of the number of nodes visited
	result := 0 // Variable to store the kth largest element

	// Call the helper function to perform reverse in-order traversal
	reverseInOrder(root, k, &count, &result)

	// Return the kth largest element
	return result
}

// reverseInOrder performs a reverse in-order traversal (right -> root -> left)
// to find the kth largest element in the BST
func reverseInOrder(node *TreeNode, k int, count *int, result *int) {
	// Base case: if the node is nil or we've already found the kth element, return
	if node == nil || *count >= k {
		return
	}

	// Step 1: Traverse the right subtree first (since we want descending order)
	reverseInOrder(node.Right, k, count, result)

	// If we've already found the kth element, stop further traversal
	if *count >= k {
		return
	}

	// Step 2: Process the current node
	*count++ // Increment the counter for each node visited
	if *count == k {
		// If the counter equals k, we've found the kth largest element
		*result = node.Val
		return
	}

	// Step 3: Traverse the left subtree (if needed)
	reverseInOrder(node.Left, k, count, result)
}
