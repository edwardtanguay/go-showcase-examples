package main

import "fmt"

func add(a, b int) int {
	return a + b
}

func addAll(values ...int) int {
	total := 0
	for _, value := range values {
		total += value
	}
	return total
}

func main() {
	fmt.Println(add(5, 2))
}
