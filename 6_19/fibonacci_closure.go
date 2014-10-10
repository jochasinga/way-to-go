package main

import "fmt"

func main() {
	var fib = Fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(fib())
	}
}

func Fibonacci() func() int {
	var n int
	return func() (result int) {
		if n == 0 {
			result = 0
		} else if n == 1 {
			result = 1
		} else {
			result = (n-1) + (n-2)
		}
		n++; return
	}
}
