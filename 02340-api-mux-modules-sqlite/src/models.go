package src

import (
	"database/sql"
	"fmt"
)

func GetEmployees(db *sql.DB) ([]Employee, error) {
	rows, err := db.Query("SELECT EmployeeID as Id, FirstName, LastName, Notes FROM Employees")
	if err != nil {
		return nil, fmt.Errorf("error querying database: %v", err)
	}
	defer rows.Close()

	var employees []Employee
	for rows.Next() {
		var e Employee
		rows.Scan(&e.Id, &e.FirstName, &e.LastName, &e.Notes)
		employee := Employee{Id: e.Id, FirstName: e.FirstName, LastName: e.LastName, Notes: e.Notes}
		employees = append(employees, employee)
	}

	return employees, nil
}

func (app *App) GetSingleEmployee(id int) (Employee, error) {
	var e Employee
	err := app.DB.QueryRow("SELECT EmployeeID as Id, FirstName, LastName, Notes FROM Employees WHERE EmployeeID = ?", id).Scan(&e.Id, &e.FirstName, &e.LastName, &e.Notes)
	if err != nil {
		if err == sql.ErrNoRows {
			return e, fmt.Errorf("no employee found with id %d", id)
		}
		return e, fmt.Errorf("error querying database: %v", err)
	}
	return e, nil
}
