package main

import "math"

type Node struct {
	Value int
	Left  *Node
	Right *Node
}

func NewNode(value int) *Node {
	return &Node{
		Value: value,
		Left:  nil,
		Right: nil,
	}
}

func IsBST(root *Node) bool {
	return isBSTHelper(root, math.MinInt64, math.MaxInt64)
}

func isBSTHelper(node *Node, min, max int) bool {
	if node == nil {
		return true
	}

	if node.Value <= min || node.Value >= max {
		return false
	}

	if !isBSTHelper(node.Left, min, node.Value) {
		return false
	}

	if !isBSTHelper(node.Right, node.Value, max) {
		return false
	}

	return true
}
