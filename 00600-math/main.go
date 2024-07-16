package main

import "fmt"
import "math"

func main() {
	
	separator()

	point1, point2, point3 := 1,2,3
	fmt.Printf("sum = %d\n", point1 + point2 + point3)

	separator()

	println(math.Pi)
	fmt.Printf("%v is of type %T", math.Pi, math.Pi)
	ceil := math.Ceil(math.Pi)
	floor := math.Floor(math.Pi)
	println(ceil)
	println(floor)
	fmt.Printf("%v is of type %T\n", ceil, ceil)
	fmt.Printf("%v is of type %T\n", floor, floor)
	fmt.Println(ceil)
	fmt.Println(floor)

	separator("mod")

	fmt.Println("nnn")

}

