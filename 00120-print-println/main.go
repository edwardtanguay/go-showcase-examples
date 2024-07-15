package main

import "fmt"

func main() {

	// for debugging
	a := 1
	b := 2
	c := 3
	println("Values: ", a, b, c)
	// println("Values: %v", c) // not possible

	fmt.Print("Hello ")
	fmt.Println("World") // with new-line
	fmt.Printf("Values: (%v)", c)

	items := []int{1, 2, 3, 4, 5}
	fmt.Println(items)
	length, err := fmt.Println(items)
	fmt.Println(length, err)
	length2, err2 := fmt.Print(items)
	fmt.Println(length2, err2)

}
