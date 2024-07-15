package main

import "fmt"

type circle struct {
	radius int
	border int
}

func main() {
	age := 34
	length := 2.322

	fmt.Printf("Age in DECIMAL is %d, Length is %f\n", age, length)
	fmt.Printf("Age in HEX is %x\n", age)

	// boolean
	fmt.Printf("Is old enough: %t", age > 10)

	// %e is exponential
	fmt.Printf("Age in DECIMAL is %d, Length in EXPONENTIAL is %e\n", age, length)

	// change order
	fmt.Printf("length is %[2]f, and age is %[1]d\n", age, length) 


	// structs
	c := circle {
		radius: 34,
		border: 2,
	}
	fmt.Printf("Radius is: %d\n", c.radius)
	fmt.Printf("c is: %v\n", c)
	fmt.Printf("c is: %+v\n", c)
	fmt.Printf("c is type: %T\n", c)

	// product strings
	message := fmt.Sprintf("length is %[2]f, and age is %[1]d", age, length) 
	fmt.Printf("The message is \"%s\".\n", message)

}
