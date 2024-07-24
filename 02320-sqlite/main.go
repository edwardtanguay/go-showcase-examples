package main

import (
	"fmt"
)

type Employee struct {
	FirstName string
	LastName  string
}

func main() {
	employees := getEmployees()
	fmt.Printf("There are %d employees:\n", len(employees))
	for i, emp := range employees {
		fmt.Printf("%d. %s\n", i, emp.FirstName+" "+emp.LastName)
	}
}
