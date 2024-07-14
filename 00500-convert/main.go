package main

import (
	"fmt"
	"math"
)

func main() {

	age := 34
	exactYears := 3.4

	newAge := float64(age) + exactYears
	fmt.Printf("we added %d and %f and got %f\n", age, exactYears, newAge)
	fmt.Printf("we added %v and %v and got %v\n", age, exactYears, newAge)
	fmt.Printf("we added %v years which rounded is %v years", exactYears, math.Round(exactYears))
}
