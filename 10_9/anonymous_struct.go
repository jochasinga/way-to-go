package main

import "fmt"

type A struct {
	f float32
	int
	string
}

func main() {
	a := A{1.20, 3, "mommy"}
	fmt.Println(a)
}
