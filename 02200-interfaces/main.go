
package main

import "fmt"
import "math"

interface Shape {
	getArea() float64
}

type Circle struct {
	radius float64
}

type Rectangle struct {
	height float64
	width float64
}

(c Circle) {
	func c.getArea() {
		return radius * math.pi
	}
} 

(r Rectangle) {
	func r.getArea() {
		return r.height * r.width
	}
}

func main() {
	c1 := Shape(2.11)
	r1 := Rectable(3.1, 4.5)

	shapes := []Shape{c1, r1}

	for _,shape := range shapes {
		fmt.Printf("Area is %f", shape.getArea())
	}
}

