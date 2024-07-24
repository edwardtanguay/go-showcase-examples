package base

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func (app *App) handleHomeRoute(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, `
<h1>Northwind API</h1>
<ul>
	<li>get all employees: <a href="http://localhost:9003/api/employees">http://localhost:9003/api/employees</a></li>
	<li>get one employee: <a href="http://localhost:9003/api/employees/1">http://localhost:9003/api/employees/1</a></li>
</ul>
	`)
}

func (app *App) handleGetEmployeesRoute(w http.ResponseWriter, _ *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	employees := []string{"emp1", "emp2"}
	json.NewEncoder(w).Encode(employees)
}

func (app *App) handleGetSingleEmployeeRoute(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r) 
	id := params["id"]
	w.Header().Set("Content-Type", "application/json")
	employees := []string{"employee ID=" + string(id)}
	json.NewEncoder(w).Encode(employees)
}