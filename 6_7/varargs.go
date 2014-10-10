package main

import "fmt"

func main() {
	f(10, 20, 43, 54, 100, 3)
}

func f(args ...int) {
	for _, v := range args {
		fmt.Println(v)
	}
}
