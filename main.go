package main

import "fmt"

func reverseString(input string) string {
	runes := []rune(input)
	n := len(runes)
	for i := 0; i < n/2; i++ {
		runes[i], runes[n-i-1] = runes[n-i-1], runes[i]
	}
	return string(runes)
}

func main() {
	input := "assignment day 19 - git"
	fmt.Println(reverseString(input))
}
