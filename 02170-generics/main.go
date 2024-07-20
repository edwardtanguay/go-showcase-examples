
package main

import "fmt"

func humanAge(dogAge int) int {
	return dogAge * 7;
}

func main() {
	dogAge := 3
	fmt.Printf("The dog is %d which in human age is %d.\n", dogAge, humanAge(dogAge))
}
