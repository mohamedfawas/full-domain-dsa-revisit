package main

import "fmt"

func reverseString(s string) string {
	if len(s) <= 1 {
		return s
	}

	return reverseString(s[1:]) + s[0:1]
}

func main() {
	str := "fawas"
	result := reverseString(str)
	fmt.Println(result)
}
