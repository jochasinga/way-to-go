package main

import (
	"fmt"
	"unicode/utf8"
)

// Count total bytes in a string
func main() {
	fmt.Println(countByte("asSASA ddd dsjkdsjs dk"))
	fmt.Println(countByte("asSASA ddd dsjkdsjsこん dk"))
}

func countByte(s string) int {
	runeCount := utf8.RuneCountInString(s)
	byteCount := len(s) - runeCount

	return byteCount + runeCount*4
}
