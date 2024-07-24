package base

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/gorilla/mux"
)

func (app *App) handleHomeRoute(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, `
<h1>Northwind Site</h1>
<hr>
<a href="/">Home</a> | <a href="/employees">Employees</a>
<hr>
<p>Welcome</p>
	`)
}

func (app *App) handleGetEmployeesRoute(w http.ResponseWriter, _ *http.Request) {
	employees, err := app.GetEmployees()
	if err != nil {
		fmt.Printf("%#v\n", err)
	} else {
		var sb strings.Builder
		fmt.Fprintf(&sb, "<h1>There are %d employees:</h1>", len(employees))
		for i, emp := range employees {
			fmt.Fprintf(&sb, "<li>(%d) %s</li>", i+1, emp.FirstName+" "+emp.LastName)
		}
		fmt.Fprint(w, sb.String())
	}
}

func (app *App) handleGetSingleEmployeeRoute(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r) 
	idStr := params["id"]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid employee ID", http.StatusBadRequest)
		return
	}
	employees, err := app.GetSingleEmployee(id)
	if err != nil {
		fmt.Printf("%#v\n", err)
	} else {
		var sb strings.Builder
		fmt.Fprintf(&sb, "<h1>There are %d employees:</h1>", len(employees))
		for i, emp := range employees {
			fmt.Fprintf(&sb, "<li>(%d) %s</li>", i+1, emp.FirstName+" "+emp.LastName)
		}
		fmt.Fprint(w, sb.String())
	}
}