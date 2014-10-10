package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(MySqrt(-9))
}

func MySqrt(arg float64) (result float64) {
	defer func() {
		errStr := recover()
		fmt.Println(errStr)
	}()
	result = math.Sqrt(arg)
	if math.IsNaN(result) {
		panic("Result is a negative number")
	}
	return
}
