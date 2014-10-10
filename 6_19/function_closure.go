package main

import "fmt"

func main() {
	var f = Adder()
	fmt.Print(f(1), " - ")
	fmt.Print(f(20), " - ")
	fmt.Print(f(300), "\n")

	var fib = Fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(fib())
	}
}

func Adder() func(int) int {
	var x int
	return func(delta int) int {
		x += delta
		return x
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
