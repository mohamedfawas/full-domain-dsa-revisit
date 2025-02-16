package main

type Node struct {
	isEnd    bool           // Marks if a word ends at this node
	children map[rune]*Node // Maps each character to its child node
}

type WordDictionary struct {
	root *Node // Root of the trie
}

// NewNode creates and initializes a new trie node
func NewNode() *Node {
	return &Node{
		isEnd:    false,                // Initially, this is not the end of any word
		children: make(map[rune]*Node), // Empty map to store children
	}
}

// Constructor initializes a new WordDictionary with an empty trie
func Constructor() WordDictionary {
	return WordDictionary{
		root: NewNode(), // Create the root node
	}
}

// AddWord adds a word to the dictionary
// Example: Adding "bad"
// 1. Start at root
// 2. Add 'b' as child of root
// 3. Add 'a' as child of 'b'
// 4. Add 'd' as child of 'a'
// 5. Mark 'd' node as end of word
func (this *WordDictionary) AddWord(word string) {
	current := this.root

	// Traverse through each character in the word
	for _, char := range word {
		// If the character doesn't exist in current node's children,
		// create a new node for it
		// Example: When adding "bad", if 'b' doesn't exist under root,
		// create a new node for 'b'
		if _, exists := current.children[char]; !exists {
			current.children[char] = NewNode()
		}
		// Move to the child node
		// Example: After adding 'b', move to the 'b' node
		current = current.children[char]
	}

	// Mark the end of the word
	// Example: After adding all chars in "bad", mark the 'd' node as end
	current.isEnd = true
}

func (this *WordDictionary) Search(word string) bool {
	wordRunes := []rune(word)

	// Define searchHelper as a nested function within Search
	var searchHelper func(node *Node, index int) bool
	searchHelper = func(node *Node, index int) bool {
		// Base case: if we've reached the end of the word
		if index == len(wordRunes) {
			return node.isEnd
		}

		char := wordRunes[index]

		// If the current character is a dot (wildcard)
		if char == '.' {
			// Try all possible children
			for _, child := range node.children {
				// If any path leads to a match, return true
				if searchHelper(child, index+1) {
					return true
				}
			}
			return false
		} else {
			// Regular character case
			child, exists := node.children[char]
			if !exists {
				return false
			}
			// Continue searching with the next character
			return searchHelper(child, index+1)
		}
	}

	// Start the search at the root
	return searchHelper(this.root, 0)
}

/**
 * Extended Example:
 *
 * 1. Add words: "bad", "dad", "mad"
 *
 * Trie structure:
 *       root
 *      / | \
 *     b  d  m
 *     |  |  |
 *     a  a  a
 *     |  |  |
 *     d  d  d
 *     *  *  * (* means isEnd=true)
 *
 * 2. Search examples:
 *
 *   a) search("bad")
 *      - Start at root
 *      - Look for 'b' in children -> found, move to 'b' node
 *      - Look for 'a' in children -> found, move to 'a' node
 *      - Look for 'd' in children -> found, move to 'd' node
 *      - Reached end of word, check if isEnd=true -> yes, return true
 *
 *   b) search(".ad")
 *      - Start at root
 *      - '.' means try all children: 'b', 'd', 'm'
 *      - For 'b': check if path "ad" exists -> yes, and ends in isEnd=true
 *      - Return true
 *
 *   c) search("b..")
 *      - Start at root
 *      - Look for 'b' in children -> found, move to 'b' node
 *      - '.' means try all children of 'b': only 'a'
 *      - For 'a': next char is '.', try all children of 'a': only 'd'
 *      - For 'd': check if isEnd=true -> yes, return true
 *
 *   d) search("bae")
 *      - Start at root
 *      - Look for 'b' in children -> found, move to 'b' node
 *      - Look for 'a' in children -> found, move to 'a' node
 *      - Look for 'e' in children -> not found, return false
 */
