package main

import "fmt"

func main() {
	fmt.Println(fibonacci(5))
}

func fibonacci(n int) func() ([]int) {

	if n <= 1 {
		res := 1
	} else {
		res := fibonacci(n-1) + fibonacci(n-2)
		return res
	}

	return func(res int) ([]int) {
		
	}
	
	for i, _ := range res {
		res[i] = (n-1) + (n-2)
	}
	return
}
