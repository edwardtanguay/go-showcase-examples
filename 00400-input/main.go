package main

import (
	"fmt"
	"bufio"
	"os"
	"strings"
	"strconv"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("What is your name? ")
	username, _ := reader.ReadString('\n')
	username = strings.TrimSpace(username)
	fmt.Printf("Your name is \"%s\".\n", username)

	fmt.Print("What is your age? ")
	ageString, _ := reader.ReadString('\n')
	ageFloat, err := strconv.ParseFloat(strings.TrimSpace(ageString), 64)
	if(err != nil) {
		fmt.Println("Sorry, there was an error:")
		fmt.Println(err)
	} else {
		fmt.Printf(`Your age to 2 decimal places is %.2f.`, ageFloat)
	}
}
