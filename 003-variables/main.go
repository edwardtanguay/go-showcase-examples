package main

import "fmt"

func separator() {
	fmt.Println("---")
}

func main() {
	var message string = "greetings"
	var age = 42
	var answer int

	fmt.Printf("The message is \"%s\".\n", message)
	fmt.Printf("The message type is \"%T\".\n", message)
	separator()
	fmt.Printf("The age is %d.\n", age)
	fmt.Printf("The age type is \"%T\".\n", age)
	separator()
	fmt.Printf("The answer is %d.\n", answer)
	fmt.Printf("The answer type is \"%T\".\n", answer)
}
