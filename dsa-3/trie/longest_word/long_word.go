package main

type Node struct {
	IsEnd    bool
	Children map[rune]*Node
}

type Trie struct {
	Root *Node
}

func NewNode() *Node {
	return &Node{
		IsEnd:    false,
		Children: make(map[rune]*Node),
	}
}

func NewTrie() *Trie {
	return &Trie{
		Root: NewNode(),
	}
}

func (t *Trie) Insert(word string) {
	current := t.Root

	for _, char := range word {
		if _, exists := current.Children[char]; !exists {
			current.Children[char] = NewNode()
		}
		current = current.Children[char]
	}

	current.IsEnd = true
}

func (t *Trie) Search(word string) bool {
	current := t.Root

	for _, char := range word {
		if _, exists := current.Children[char]; !exists {
			return false
		}
		current = current.Children[char]
	}

	return current.IsEnd
}

func (t *Trie) Prefix(word string) bool {

	current := t.Root
	for _, char := range word {
		if _, exists := current.Children[char]; !exists {
			return false
		}
		current = current.Children[char]
	}

	return true
}
