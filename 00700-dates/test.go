package main

import (
	"fmt"
	"math/rand"
	"time"
)

func addCode(idCode string) string {
	rand.Seed(time.Now().UnixNano())
	randomNumber := rand.Intn(883) + 1
	if randomNumber == 1 {
		panic("there was a system error, execution stopped")
	}
	return idCode + "-8273843"
}

func separator() {
	fmt.Println("---")
}

func main() {

	now := time.Now()
	launchDateTime := time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC)

	fmt.Println(addCode("book"))
	separator()
	fmt.Printf("Time is %v\n", now)
	fmt.Printf("Go launched at %v\n", launchDateTime)
	fmt.Printf("Go launched at %v\n", launchDateTime.Format(time.ANSIC))
	fmt.Printf("Go launched at %v\n", launchDateTime.Format(time.RFC1123))
	separator()
	t := time.Now()
	fmt.Print(t.Format("2006-01-02"))
}
