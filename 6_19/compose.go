package main

import "fmt"
import "strings"

func main() {
	addJpeg := MakeAddSuffix(".jpg")
	addBmp := MakeAddSuffix(".bmp")
	fmt.Println(addJpeg("file"))
	fmt.Println(addBmp("file"))
}

func MakeAddSuffix(suffix string) func(string) string {
	return func(name string) string {
		if !strings.HasSuffix(name, suffix) {
			return name + suffix
		}
		return name
	}
}