package main

import "fmt"
import "math/rand"

func main() {
	message := "hello"
	randomNumber := rand.Intn(100)
	fmt.Printf("The message is \"%s\" and the random number is %d.",message, randomNumber)
}