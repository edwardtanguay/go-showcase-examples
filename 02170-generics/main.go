
package main

import "fmt"

func humanAge [T int|float64](dogAge T) T {
	return dogAge * 7;
}

func main() {
	dogAges := []int{12,4,2,6,1}

	for _, dogAge := range dogAges {
		humanAge := humanAge((dogAge))
		fmt.Printf("The dog is %d (type %T) which in human age is %d (type %T).\n", dogAge, dogAge, humanAge,humanAge)
	}

	exactDogAge := 4.32
	exactHumanAge := humanAge((exactDogAge))
	fmt.Printf("The exact age of the dog is %.2f (type %T)and the exact human age is %.2f (type %T).", exactDogAge, exactDogAge, exactHumanAge, exactHumanAge)
}
