
package main

import "fmt"

func main() {
	separator()
	for i := 1; i <= 15; i++ {
		fmt.Printf("when a dog is %d years old, it's human age is %d\n", i, i * 7)
	}

	separator()
	for i := 1; i <= 15; i++ {
		fmt.Printf("Dog: %3d   Human: %4d\n", i, i*7)
	}


}
