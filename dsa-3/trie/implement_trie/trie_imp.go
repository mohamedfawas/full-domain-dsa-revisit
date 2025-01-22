package main

import "fmt"

type Node struct {
	IsEnd    bool
	Children map[rune]*Node
}

func NewNode() *Node {
	return &Node{
		IsEnd:    false,
		Children: make(map[rune]*Node),
	}
}

type Trie struct {
	Root *Node
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

func (t *Trie) StartsWith(word string) bool {
	current := t.Root

	for _, char := range word {
		if _, exists := current.Children[char]; !exists {
			return false
		}
		current = current.Children[char]
	}

	return true
}

func main() {
	// Create a new empty Trie
	trie := NewTrie()

	// Add some words
	trie.Insert("hello")
	trie.Insert("world")
	trie.Insert("hi")

	// Try searching for words
	fmt.Println("Searching for 'hello':", trie.Search("hello")) // true
	fmt.Println("Searching for 'world':", trie.Search("world")) // true
	fmt.Println("Searching for 'hi':", trie.Search("hi"))       // true
	fmt.Println("Searching for 'hey':", trie.Search("hey"))     // false

	// Try prefix searches
	fmt.Println("Prefix 'he':", trie.StartsWith("he")) // true
	fmt.Println("Prefix 'wo':", trie.StartsWith("wo")) // true
	fmt.Println("Prefix 'ha':", trie.StartsWith("ha")) // false
}
