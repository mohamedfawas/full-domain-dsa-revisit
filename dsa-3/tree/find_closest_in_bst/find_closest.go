package main

import "fmt"

type Node struct {
	Value int
	Left  *Node
	Right *Node
}

type BST struct {
	Root *Node
}

func (bst *BST) FindClosest(target int) int {
	if bst.Root == nil {
		return 0
	}

	return findClosestHelper(bst.Root, target, bst.Root.Value)
}

func findClosestHelper(node *Node, target, closest int) int {
	if node == nil {
		return closest
	}

	currentDiff := abs(target - node.Value)
	closestDiff := abs(target - closest)

	if currentDiff < closestDiff {
		closest = node.Value
	}

	if target < node.Value {
		return findClosestHelper(node.Left, target, closest)
	} else {
		return findClosestHelper(node.Right, target, closest)
	}
}

func abs(x int) int {
	if x < 0 {
		return -x
	}

	return x
}

func main() {
	// Create a sample tree:
	//       10
	//      /  \
	//     5    15
	//    / \   / \
	//   2   7 13  22

	bst := &BST{
		Root: &Node{
			Value: 10,
			Left: &Node{
				Value: 5,
				Left:  &Node{Value: 2},
				Right: &Node{Value: 7},
			},
			Right: &Node{
				Value: 15,
				Left:  &Node{Value: 13},
				Right: &Node{Value: 22},
			},
		},
	}

	result := bst.FindClosest(12)
	fmt.Printf("The closest value to 12 is: %d\n", result)
}
