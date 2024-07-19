
package main

import "fmt"

func getHumanAge(dogAge int) int {
	return dogAge * 7
}

func main() {
	dogAge := 5
	fmt.Printf("Dog age is %d which in human years is %d.\n",dogAge, getHumanAge(dogAge))
}
