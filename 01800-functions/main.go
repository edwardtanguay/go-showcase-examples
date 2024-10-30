package main

import "fmt"

// types can be set for all parameters
func add(a, b int) int {
	return a + b
}

// unknown number of arguments
func addAll(values ...int) int {
	total := 0
	for _, value := range values {
		total += value
	}
	return total
}

// multiple return types
func getInfo(values ...int) (int, int) {
	total := 0
	for _, value := range values {
		total += value
	}
	return total, len(values)
}

func main() {
	separator()
	fmt.Println(add(5, 2))
	separator()
	fmt.Println(addAll(1, 2, 3, 4, 5))
	separator()
	sum, count := getInfo(1,2,3,4,5,6,7,8,9)
	fmt.Printf("There are %v numbers that add up to %v\n", count, sum)

	fmt.Printf("The answer is %v.\n", subtract(8,7))
}
