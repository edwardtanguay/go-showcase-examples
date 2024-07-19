package main

import "fmt"

func main() {
	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
	}
	word := "dictionary"
	fmt.Printf("The sum is %d\n", sum)
	fmt.Printf("The length of the word \"%s\" is %d\n", word, len(word))
}
