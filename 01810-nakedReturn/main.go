
package main

import "fmt"

func getHumanAge(dogAge int) (humanAge int) {
	humanAge = dogAge * 7
	return
}

func main() {
	dogAge := 4
	fmt.Printf("Dog age is %d which in human years is %d.\n",dogAge, getHumanAge(dogAge))
}
