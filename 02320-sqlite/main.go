package main

import (
	"fmt"
	"tanguay.info/02320-sqlite/models"
)

func main() {
	employees := models.GetEmployees()
	fmt.Printf("There are %d employees:\n", len(employees))
	for i,emp := range employees {
		fmt.Printf("%d. %s\n", i, emp.FirstName + " " + emp.LastName)
	}
}
