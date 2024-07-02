package main

import "fmt"

func printValues(i int, p *int) {
	fmt.Printf("The age is %d\n", i)
	fmt.Printf("The pointer is %d\n", p)
	fmt.Printf("The pointer points to the value %d\n", *p)
}

func main() {
	age := 34
	pAge := &age

	printValues(age, pAge)
	*pAge += 1

	fmt.Println("--- VALUE UNDER POINTER CHANGED")
	printValues(age, pAge)
}
