package main

import (
	"fmt"
	"time"
)

func makePayment(money chan float64) {
	amount := <-money
	fmt.Printf("Received: %.2f\n", amount)
}

func main() {
	money := make(chan float64)
	go makePayment(money)

	money <- 34.99

	time.Sleep(2 * time.Second)
}
