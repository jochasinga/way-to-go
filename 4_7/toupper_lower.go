package main

import (
	"fmt"
	"strings"
)

func main() {
	orig := "Hey, how are you George?"
	var lower string
	var upper string
	fmt.Printf("the original string is: %s\n", orig)
	lower = strings.ToLower(orig)
	fmt.Printf("The lowercase string is: %s\n", lower)
	upper = strings.ToUpper(orig)
	fmt.Printf("The uppercase string is: %s\n", upper)
}
