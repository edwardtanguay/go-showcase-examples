package base

import (
	"fmt"
	"net/http"
	"strings"
)

func (app *App) handleHomeRoute(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the API.")
}

func (app *App) handleEmployeesRoute(w http.ResponseWriter, r *http.Request) {
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
