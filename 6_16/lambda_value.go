package main

import "fmt"

func main() {
	fv := func() {
		fmt.Println("Hello World")
	}
	fv()
	fmt.Printf("fv has the value and type of %v and %T\n", fv, fv)
}
