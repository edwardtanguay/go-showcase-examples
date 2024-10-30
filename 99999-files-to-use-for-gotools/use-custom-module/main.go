package main

import (
	"fmt"

	"github.com/edwardtanguay/gotools"
)

func main() {
	lines, _ := gotools.GetLinesFromFile("main.go")
	for index,line := range lines {
		fmt.Printf("%02d | %s\n", index+1, line)
	}

	for i := 1; i <= 4; i++ {
		fmt.Printf("Status: %s\n", gotools.SmartPlural(i, "card", ""))
		fmt.Printf("Status: %s\n", gotools.SmartPlural(i, "bus", "buses"))
	}

	fmt.Printf("The test data is: %s\n", gotools.GetTestData())

	employees := gotools.GetEmployees()
	fmt.Printf("Employees: %#v\n", employees)

}