package main

import (
	"fmt"
	"strconv"
)

func main() {
	var orig string = "ABC"
	an, err := strconv.Atoi(orig)
	if err != nil {
		fmt.Printf("orig %s is not an integer - exiting with error\n", orig)
		return
	}
	fmt.Printf("The integer is %d\n", an)
}
