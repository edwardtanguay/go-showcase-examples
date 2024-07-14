package main

import (
	"fmt"
	"math/rand"
)

func main() {

	// random number
	dow := rand.Intn(7) + 1
	fmt.Println("Day", dow)

	var result string
	switch dow {
	case 1:
		result = "Sunday"
	case 2:
		result = "Monday"
	case 3:
		result = "Tuesday"
	default:
		result = "(some other day)"
	}

	fmt.Println(result)
	separator()

	appState := "offline"

	if appState == "online" {
		fmt.Println("app is online")
	} else if appState == "offline" {
		fmt.Println("app is OFFLINE")
	} else {
		fmt.Println("app state uncertain")
	}

}
