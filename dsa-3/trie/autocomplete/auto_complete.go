package main

import "fmt"

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

func (t *Trie) findNode(prefix string) *Node {
	current := t.Root

	for _, char := range prefix {
		if _, exists := current.Children[char]; !exists {
			return nil
		}
		current = current.Children[char]
	}

	return current
}

func findAllWords(node *Node, prefix string, words *[]string) {
	if node.IsEnd {
		*words = append(*words, prefix)
	}

	for char, childNode := range node.Children {
		findAllWords(childNode, prefix+string(char), words)
	}
}

func (t *Trie) AutoComplete(prefix string) []string {
	node := t.findNode(prefix)

	var words []string
	findAllWords(node, prefix, &words)
	return words
}

func main() {
	// Create a new word tree
	trie := NewTrie()

	// Add some words to our tree
	words := []string{
		"cat",
		"car",
		"card",
		"carpet",
		"dog",
		"deer",
		"deal",
	}

	for _, word := range words {
		trie.Insert(word)
	}

	// Try finding words that start with "ca"
	fmt.Println("Words starting with 'ca':")
	fmt.Println(trie.AutoComplete("ca")) // Will print: [cat car card carpet]

	// Try finding words that start with "de"
	fmt.Println("\nWords starting with 'de':")
	fmt.Println(trie.AutoComplete("de")) // Will print: [deer deal]
}
