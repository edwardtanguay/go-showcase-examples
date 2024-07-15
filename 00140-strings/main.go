package main

import (
	"fmt"
	"strings"
)

func main() {

	// length
	title := "LAMP"
	fmt.Printf("Length is %d\n", len(title))

	// iterate over string
	for i, ch := range title {
		fmt.Printf("%d: %c\n", i, ch)
	}

	message := "go get the book and bring it back"
	fmt.Printf("book: %t\n", strings.Contains(message, "book"))
	fmt.Printf("tiger: %t\n", strings.Contains(message, "tiger"))
	fmt.Printf("vowels: %t\n", strings.ContainsAny(message, "aeiou"))
	fmt.Printf("index of book: %d\n", strings.Index(message, "book"))
	fmt.Printf("index of tiger: %d\n", strings.Index(message, "tiger"))
	fmt.Printf("first index of b: %d\n", strings.Index(message, "b"))
	fmt.Printf("last index of b: %d\n", strings.LastIndex(message, "b"))

	fileName := "test.txt"
	fmt.Printf("starts with: %t\n", strings.HasPrefix(fileName, "test"))
	fmt.Printf("ends with: %t\n", strings.HasSuffix(fileName, ".txt"))
	fmt.Printf("ends with: %t\n", strings.HasSuffix(fileName, ".gif"))
	fmt.Printf("number of t's: %d\n", strings.Count(message, "t"))
	fmt.Printf("number of z's: %d\n", strings.Count(message, "z"))

	// count vowels
	vowels := "aeiou"
	for _, ch := range vowels {
		fmt.Printf("number of \"%c\" = %d\n", ch, strings.Count(message, string(ch)))
	}

}
