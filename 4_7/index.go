package main

import (
	"fmt"
	"strings"
)

func main() {
	str := "This is an example of a string"
	fmt.Printf("Where is the string \"%s\" in %s? ", str, "a")
	fmt.Printf("%d\n", strings.Index(str, "a"))
	fmt.Printf("Where is the last character of string \"%s\" in %s? ", str, "a")
	fmt.Printf("%d\n", strings.LastIndex(str, "a"))
}
