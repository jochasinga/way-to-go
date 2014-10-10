package main

import "fmt"

func main() {
	defer func () {
		str := recover()
		fmt.Println(str)
	}()

	var p *int = nil
	*p = 0
}
