package main

import "fmt"

func reverseString(s string) {
	for _, char := range s {
		defer fmt.Printf("%c", char)
	}
}

func main() {
	reverseString("fawas")
}
