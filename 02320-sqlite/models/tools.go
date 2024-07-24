package models 

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

func GetEmployees() (employees []Employee) {
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
		employee := Employee{FirstName: e.FirstName, LastName: e.LastName}
		employees = append(employees, employee)
	}

	return
}
