package main

import (
	"fmt"
	"unsafe"
)

type File struct {
	fd   int    // file descriptor number
	name string // file name
}

func NewFile(fd int, name string) *File {
	if fd < 0 {
		return nil
	}
	return &File{fd, name}
}

func main() {
	f := NewFile(10, "./test.txt")
	fmt.Println(f)
	size := unsafe.Sizeof(f)
	fmt.Println(size)
}
