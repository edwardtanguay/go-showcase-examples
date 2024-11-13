package main

import (
	"fmt"
	"math/rand"
)

func addNumber(nums *[]int) {
	*nums = append(*nums, rand.Intn(10)+1)
}

func test001() {
	nums := []int{6, 2}
	addNumber(&nums)
	addNumber(&nums)
	addNumber(&nums)

	fmt.Println("Result:", nums)
}

func main() {
	test001()
}
