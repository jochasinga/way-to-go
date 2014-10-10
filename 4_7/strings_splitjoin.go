package main

import (
	"fmt"
	"strings"
)
func main() {
        str := "The quick brown fox jumps over the lazy dog"
	sl := strings.Fields(str)
	fmt.Printf("Splitted in slice: %v\n", sl)
	for _, v := range sl {
		fmt.Printf("%s - ", v)
	}
	fmt.Println()
	str2 := "GO1|The ABC of Go|25"
	sl2 := strings.Split(str2, "|")
	fmt.Printf("Splitted in slice %v\n", sl2)
	for _, v := range sl2 {
		fmt.Printf("%s - ", v)
	}
	str3 := strings.Join(sl2, ";")
	fmt.Printf("sl2 joined by ;: %s\n", str3)
}