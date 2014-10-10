package main

import "fmt"

func main() {
	fmt.Println(mat("multiply", 2, 3))
}

func mat(choice string, a, b int) (result int) {

	switch choice {
	case "add":
		result = a + b
	case "multiply":
		result = a * b
	case "subtract":
		result = a - b
	default:
		result = a + b
	}
	return
}
