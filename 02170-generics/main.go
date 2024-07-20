
package main

import "fmt"

func humanAge(dogAge int) int {
	return dogAge * 7;
}

func main() {
	dogAges := []int{12,4,2,6,1}

	for _, dogAge := range dogAges {
		fmt.Printf("The dog is %d which in human age is %d.\n", dogAge, humanAge(dogAge))
	}
}
