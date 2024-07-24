package main

import (
	"tanguay.info/02320-sqlite/models"
)

func main() {
	// employees, err := models.GetEmployees()
	a := models.App{}
	a.Port = 9003
	a.Initialize()
	a.Run()
	defer a.DB.Close()
	// if err != nil {
	// 	fmt.Printf("Cannot get employees: %v\n", err)
	// } else {
	// 	fmt.Printf("There are %d employees:\n", len(employees))
	// }
}
