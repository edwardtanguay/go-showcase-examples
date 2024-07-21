package main

import (
	"fmt"
	"time"
)

func makePayment(money chan float64) {
	select {
	case amount := <-money:
		fmt.Printf("Received: %.2f\n", amount)
	case <-time.After(time.Second * 1):
		fmt.Print("NOTHING RECEIVED")
	}
}

func main() {
	money := make(chan float64)
	go makePayment(money)

	money <- 34.99 // comment out for nothing to be received

	time.Sleep(2 * time.Second)
}
