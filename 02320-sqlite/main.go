package main

import (
	"fmt"

	"tanguay.info/02320-sqlite/models"
)

func main() {
	employees, err := models.GetEmployees()
	if err != nil {
		fmt.Printf("Cannot get employees: %v\n", err)
	} else {
		fmt.Printf("There are %d employees:\n", len(employees))
		for i, emp := range employees {
			fmt.Printf("%d. %s\n", i+1, emp.FirstName+" "+emp.LastName)
		}
	}
}
