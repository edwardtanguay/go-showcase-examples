
package main

import "fmt"

func PrintAnything(v interface{}) {
	fmt.Printf("The variable of type %T has a value of: %v\n", v, v)
}

func main() {
	PrintAnything("hello")
	PrintAnything(123)
	PrintAnything(1.34)
	PrintAnything(true)
}
