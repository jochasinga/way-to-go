package main

import "fmt"

// this function changes reply:
func Multiply(a, b int, reply *int) {
	*reply = a * b
}

func main() {
	n := 0
	rep := &n
	Multiply(10, 5, rep)
	fmt.Println("Multiply:", *rep) // Multiply: 50
}
