package main

type Node struct {
	isEnd    bool
	children map[rune]*Node
}

func NewNode() *Node {
	return &Node{
		isEnd:    false,
		children: make(map[rune]*Node),
	}
}

type Trie struct {
	root *Node
}

func Constructor() Trie {
	return Trie{
		root: NewNode(),
	}
}

func (this *Trie) Insert(word string) {
	current := this.root

	for _, char := range word {
		if _, exists := current.children[char]; !exists {
			current.children[char] = NewNode()
		}
		current = current.children[char]
	}

	current.isEnd = true
}

func (this *Trie) Search(word string) bool {
	current := this.root

	for _, char := range word {
		if _, exists := current.children[char]; !exists {
			return false
		}
		current = current.children[char]
	}

	return current.isEnd
}

func (this *Trie) StartsWith(prefix string) bool {
	current := this.root

	for _, char := range prefix {
		if _, exists := current.children[char]; !exists {
			return false
		}
		current = current.children[char]
	}

	return true
}
