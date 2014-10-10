package main

import "fmt"

type Rectangle struct {
	length float64
	width  float64
}

func (r *Rectangle) Area() float64 {
	return r.length * r.width
}

func (r *Rectangle) Perimeter() float64 {
	return (r.length + r.width) * 2
}

func main() {
	//rect := new(Rectangle)
	rect := Rectangle{10.25, 3.0}
	fmt.Println(rect.Area())
	fmt.Println(rect.Perimeter())
}
