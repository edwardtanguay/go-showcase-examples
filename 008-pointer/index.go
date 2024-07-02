package main

import "fmt"

func main() {
	age := 34
	pAge := &age

	fmt.Printf("The age is %d\n", age)
	fmt.Printf("The pointer is %d\n", pAge)
	fmt.Printf("The pointer points to the value %d\n", *pAge)

	*pAge += 1

	fmt.Println("--- VALUE UNDER POINTER CHANGED")
	fmt.Printf("The age is %d\n", age)
	fmt.Printf("The pointer is %d\n", pAge)
	fmt.Printf("The pointer points to the value %d\n", *pAge)
}