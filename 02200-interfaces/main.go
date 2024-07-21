package main

import (
	"fmt"
	"math"
)

type Shape interface {
	getArea() float64
}

type Circle struct {
	Radius float64
}

func (c Circle) getArea() float64 {
	return c.Radius * c.Radius * math.Pi
}

type Rectangle struct {
	Height float64
	Width float64
}

func (r Rectangle) getArea() float64 {
	return r.Height * r.Width
}

func main() {
	c1 := Circle{3.21}
	r1 := Rectangle{3.23, 2.23}

	shapes := []Shape{c1, r1}

	for _, shape := range shapes {
		fmt.Printf("Shape of type %T has an area of %f.\n", shape, shape.getArea())
	}
}

