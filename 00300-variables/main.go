package main

import "fmt"

const path = "usr/bin"

func separator() {
	fmt.Println("---")
}

func main() {
	var message string = "greetings"
	var age = 42
	var answer int
	rank := 5.4 // same as with var, but only inside functions

	fmt.Printf("The message is \"%s\".\n", message)
	fmt.Printf("The message type is \"%T\".\n", message)
	separator()
	fmt.Printf("The age is %d.\n", age)
	fmt.Printf("The age type is \"%T\".\n", age)
	separator()
	fmt.Printf("The answer is %d.\n", answer)
	fmt.Printf("The answer type is \"%T\".\n", answer)
	separator()
	fmt.Printf("The rank is %f.\n", rank)
	fmt.Printf("The rank is %.2f.\n", rank)
	fmt.Printf("The rank type is \"%T\".\n", rank)
	separator()
	fmt.Printf("The constant path is \"%s\".\n", path)
	fmt.Printf("The constant path type is \"%T\".\n", path)
}
