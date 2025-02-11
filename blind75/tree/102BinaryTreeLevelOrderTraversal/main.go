package main

// First, we define the TreeNode structure that represents each node in our binary tree
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrder(root *TreeNode) [][]int {
	// If the tree is empty (root is nil), return an empty slice
	if root == nil {
		return [][]int{}
	}

	// Initialize result slice to store the final level-order traversal
	// This will be a slice of slices, where each inner slice represents one level
	// Example: [[3],[9,20],[15,7]] for the tree:
	//     3
	//    / \
	//   9  20
	//      / \
	//    15   7
	result := [][]int{}

	// Initialize a queue to help us process nodes level by level
	// Queue will store pointers to TreeNode
	queue := []*TreeNode{root}

	// Continue processing while there are nodes in the queue
	for len(queue) > 0 {
		// Get the number of nodes at current level
		// This is important because we need to process all nodes
		// at the current level before moving to the next level
		levelSize := len(queue)

		// Create a slice to store values of nodes at current level
		currentLevel := []int{}

		// Process all nodes at current level
		for i := 0; i < levelSize; i++ {
			// Remove the first node from queue (dequeue operation)
			// In Go, this is done by slicing the queue
			node := queue[0]
			queue = queue[1:]

			// Add the current node's value to currentLevel slice
			currentLevel = append(currentLevel, node.Val)

			// Add left child to queue if it exists
			// This child will be processed in the next level
			if node.Left != nil {
				queue = append(queue, node.Left)
			}

			// Add right child to queue if it exists
			// This child will be processed in the next level
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}

		// After processing all nodes at current level,
		// append currentLevel slice to result
		result = append(result, currentLevel)
	}

	return result
}

/*
Example walkthrough of the algorithm:

Consider this tree:
     3
    / \
   9  20
      / \
    15   7

1. Initial state:
   - queue = [3]
   - result = []

2. First iteration (Level 0):
   - levelSize = 1 (only root node)
   - Process node 3
   - Add its children (9, 20) to queue
   - currentLevel = [3]
   - queue = [9, 20]
   - result = [[3]]

3. Second iteration (Level 1):
   - levelSize = 2 (nodes 9 and 20)
   - Process node 9 (no children)
   - Process node 20 (has children 15, 7)
   - Add 15 and 7 to queue
   - currentLevel = [9, 20]
   - queue = [15, 7]
   - result = [[3], [9,20]]

4. Third iteration (Level 2):
   - levelSize = 2 (nodes 15 and 7)
   - Process node 15 (no children)
   - Process node 7 (no children)
   - currentLevel = [15, 7]
   - queue = [] (empty)
   - result = [[3], [9,20], [15,7]]

5. Queue is empty, algorithm terminates
   Final result: [[3], [9,20], [15,7]]
*/
