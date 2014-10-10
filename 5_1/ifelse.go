package main

import "fmt"

func main() {
	var first int = 10
	var cond int
	if first <= 0 {
		fmt.Printf("first is less than or equal to 0\n")
	} else {
		fmt.Printf("first is 5 or greater\n")
	}

	if cond = 5; cond > 10 {
		fmt.Printf("cond is greater than 10\n")
	} else {
		fmt.Printf("cond is not greater than 10\n")
	}
}
