package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

type Employee struct {
	FirstName string
	LastName  string
}

func getEmployees() (employees []Employee) {
	db, err := sql.Open("sqlite3", "./data/northwind.sqlite")
	if err != nil {
		log.Fatal(err.Error())
	}
	defer db.Close()

	rows, err := db.Query("SELECT FirstName, LastName FROM Employees")
	if err != nil {
		log.Fatal(err.Error())
	}
	defer rows.Close()

	for rows.Next() {
		var e Employee
		rows.Scan(&e.FirstName, &e.LastName)
		employee := Employee{e.FirstName, e.LastName}
		employees = append(employees, employee)
	}

	return
}

func main() {
	employees := getEmployees()
	fmt.Printf("There are %d employees:\n", len(employees))
	for i,emp := range employees {
		fmt.Printf("%d. %s\n", i, emp.FirstName + " " + emp.LastName)
	}
}
