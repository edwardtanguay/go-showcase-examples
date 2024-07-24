package models

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	_ "github.com/mattn/go-sqlite3"
)

type App struct {
	DB *sql.DB
	Port int
}

func (a *App) Initialize() error {
	var err error
	a.DB, err = sql.Open("sqlite3", "./data/northwind.sqlite")
	if err != nil {
		return fmt.Errorf("error opening database: %v", err)
	}
	defer a.DB.Close()
	return nil
}

func handleHomeRoute(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the API.")
}

func handleEmployeesRoute(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "here the employees")
}

func (a *App) Run() {
	http.HandleFunc("/", handleHomeRoute)
	http.HandleFunc("/employees", handleEmployeesRoute)
	fmt.Printf("API running at http://localhost:%d\n", a.Port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", a.Port), nil))
}

// func GetEmployees() ([]Employee, error) {

// 	rows, err := db.Query("SELECT FirstName, LastName FROM Employees")
// 	if err != nil {
// 		return nil, fmt.Errorf("error querying database: %v", err)
// 	}
// 	defer rows.Close()

// 	var employees []Employee
// 	for rows.Next() {
// 		var e Employee
// 		rows.Scan(&e.FirstName, &e.LastName)
// 		employee := Employee{FirstName: e.FirstName, LastName: e.LastName}
// 		employees = append(employees, employee)
// 	}

// 	return employees, nil
// }
