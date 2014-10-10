package main

import (
	"fmt"
	"strings"
)

func main() {
	str := "This is an example of a string"
	fmt.Printf("T/F? Does the string \"%s\" contains %s? ", str, "ex")
	fmt.Printf("%t\n", strings.Contains(str, "ex"))
}
