package src

import "fmt"

func (app *App) GetEmployees() ([]Employee, error) {
	rows, err := app.DB.Query("SELECT EmployeeID as Id, FirstName, LastName, Notes FROM Employees")
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

func (app *App) GetSingleEmployee(id int) ([]Employee, error) {
	rows, err := app.DB.Query("SELECT EmployeeID as Id, FirstName, LastName, Notes FROM Employees WHERE EmployeeID = ?", id)
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
