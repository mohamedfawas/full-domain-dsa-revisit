package main

type Node struct {
	Value int
	Left  *Node
	Right *Node
}

type BST struct {
	Root *Node
}

func NewNode(value int) *Node {
	return &Node{
		Value: value,
		Left:  nil,
		Right: nil,
	}
}

func InorderTraversal(node *Node, elements *[]int) {
	if node == nil {
		return
	}
	InorderTraversal(node.Left, elements)
	*elements = append(*elements, node.Value)
	InorderTraversal(node.Right, elements)
}

func (bst *BST) KthLargest(k int) int {
	if bst.Root == nil {
		return -1
	}

	elements := []int{}
	InorderTraversal(bst.Root, &elements)
	n := len(elements)

	if k > n || k <= 0 {
		return -1
	}

	return elements[n-k]
}
