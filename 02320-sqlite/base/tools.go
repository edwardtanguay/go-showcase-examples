package base

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"strings"

	_ "github.com/mattn/go-sqlite3"
)

type App struct {
	DB   *sql.DB
	Port int
}

func (a *App) Initialize() error {
	var err error
	a.DB, err = sql.Open("sqlite3", "./data/northwind.sqlite")
	if err != nil {
		return fmt.Errorf("error opening database: %v", err)
	}
	return nil
}

func handleHomeRoute(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the API.")
}

func (a *App) handleEmployeesRoute(w http.ResponseWriter, r *http.Request) {
	employees, err := a.GetEmployees()
	if err != nil {
		fmt.Printf("%#v\n", err)
	} else {
		var sb strings.Builder
		fmt.Fprintf(&sb, "<h1>There are %d employees:</h1>", len(employees))
		for i, emp := range employees {
			fmt.Fprintf(&sb, "<li>nnn %d. %s</li>", i+1, emp.FirstName+" "+emp.LastName)
		}
		fmt.Fprint(w, sb.String())
	}
}

func (a *App) Run() {
	http.HandleFunc("/", handleHomeRoute)
	http.HandleFunc("/employees", a.handleEmployeesRoute)
	fmt.Printf("API running at http://localhost:%d\n", a.Port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", a.Port), nil))
}

func (a *App) GetEmployees() ([]Employee, error) {
	rows, err := a.DB.Query("SELECT FirstName, LastName FROM Employees")
	if err != nil {
		return nil, fmt.Errorf("error querying database: %v", err)
	}
	defer rows.Close()

	var employees []Employee
	for rows.Next() {
		var e Employee
		rows.Scan(&e.FirstName, &e.LastName)
		employee := Employee{FirstName: e.FirstName, LastName: e.LastName}
		employees = append(employees, employee)
	}

	return employees, nil
}
