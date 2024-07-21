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

type Rectangle struct {
	Height float64
	Width  float64
}

func (c Circle) getArea() float64 {
	return c.Radius * math.Pi
}

func (r Rectangle) getArea() float64 {
	return r.Height * r.Width
}

func main() {
	c1 := Circle{2.11}
	r1 := Rectangle{3.1, 4.5}

	shapes := []Shape{c1, r1}

	for _, shape := range shapes {
		fmt.Printf("Area for type %T is %f\n", shape, shape.getArea())
	}
}
