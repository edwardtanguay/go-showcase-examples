package main

import (
	"fmt"
	"strconv"
)

func main() {
	num := 100
	str := string(rune(num))
	fmt.Printf("%v is of type %T\n", str, str) // not "100" but 'd' as rune

	str2 := strconv.Itoa(num)
	fmt.Printf("%v is of type %T\n", str2, str2)
}
