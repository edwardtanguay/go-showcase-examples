package src

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"
	"strings"

	"github.com/gorilla/mux"
)

func (app *App) handleHomeRoute(w http.ResponseWriter, r *http.Request) {
	type Data struct {
		Header  template.HTML
		Version string
	}
	data := Data{
		Header:  template.HTML(Header()),
		Version: Config().Version,
	}
	t, _ := template.New("page").Parse(`
{{.Header}}
<p>Welcome to the Northwind site. {{.Version}}</p>
	`)
	t.Execute(w, data)
}

func (app *App) handleGetEmployeesRoute(w http.ResponseWriter, _ *http.Request) {
	employees, err := GetEmployees(app.DB)
	if err != nil {
		fmt.Printf("%#v\n", err)
	} else {
		var sb strings.Builder
		fmt.Fprintf(&sb, "<h2>There are %d employees:</h2>", len(employees))
		for _, emp := range employees {
			fmt.Fprintf(&sb, "<li><a href=\"/employees/%d\">%s</a></li>", emp.Id, emp.FirstName+" "+emp.LastName)
		}
		fmt.Fprint(w, Header()+sb.String())
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
	emp, err := GetSingleEmployee(app.DB, id)
	if err != nil {
		fmt.Printf("%#v\n", err)
	} else {
		var sb strings.Builder
		fmt.Fprintf(&sb, "<h3>%s</h3>", emp.FirstName+" "+emp.LastName)
		fmt.Fprintf(&sb, "<p>%s</p>", emp.Notes)
		fmt.Fprint(w, Header()+sb.String())
	}
}
