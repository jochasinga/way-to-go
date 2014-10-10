package main

import "fmt"

type Point2d struct {
	x, y int
}

type Point3d struct {
	Point2D
	z int
}

type Polar struct {
	radius float64
	angle int
}

func (p2d *Point2d) Abs() float64 {
	
}
