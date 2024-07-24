package base

import "fmt"

func (app *App) GetEmployees() ([]Employee, error) {
	rows, err := app.DB.Query("SELECT FirstName, LastName FROM Employees")
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

func (app *App) GetSingleEmployee(id int) ([]Employee, error) {
	rows, err := app.DB.Query("SELECT FirstName, LastName FROM Employees WHERE EmployeeID = ?", id)
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
