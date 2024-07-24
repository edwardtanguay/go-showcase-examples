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

func getEmployees() {
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
		var e Employee

		rows.Scan(&e.FirstName, &e.LastName)
		fmt.Printf("EMPLOYEE: %s, %s\n", e.LastName, e.FirstName)
	}

}

func main() {
	println("test")
	getEmployees()
}
