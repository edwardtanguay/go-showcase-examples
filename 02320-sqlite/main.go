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

	rows, err := db.Query("SELECT FirstName, LastName FROM Employees")
	if err != nil {
		log.Fatal(err.Error())
	}
	defer rows.Close()

	for rows.Next() {
		// var e Employee
		// rows.Scan(&e.FirstName, &e.LastName)
		// employee := Employee{e.FirstName, e.LastName}
		// fmt.Printf("%#v\n", employee)

		employee := Employee{"fff", "lll"}
		employees = append(employees, employee)
		return
	}
	return
}

func main() {
	employees := getEmployees()
	fmt.Printf("There are %d employees.\n", len(employees))
}
