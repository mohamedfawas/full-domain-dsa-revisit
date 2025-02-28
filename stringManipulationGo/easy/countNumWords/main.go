package main

import (
	"fmt"
	"strings"
)

func countNumWords(word string) int {
	words := strings.Fields(word)
	return len(words)
}

func main() {
	fmt.Println(countNumWords("Hello World"))
}
