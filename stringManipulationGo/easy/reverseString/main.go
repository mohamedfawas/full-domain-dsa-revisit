package main

import "fmt"

func reverseString(word string) string {
	var reverseString string
	for i := len(word) - 1; i >= 0; i-- {
		reverseString += string(word[i])
	}

	return reverseString
}

func main() {
	fmt.Println(reverseString("hello"))
	fmt.Println(reverseString("world"))

}
