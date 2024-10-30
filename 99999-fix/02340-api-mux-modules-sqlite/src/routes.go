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
		SiteBegin template.HTML
		SiteEnd template.HTML
		Header  template.HTML
		Version string
	}
	data := Data{
		SiteBegin: template.HTML(SiteBegin()),
		SiteEnd: template.HTML(SiteEnd()),
		Header:  template.HTML(Header()),
		Version: Config().Version,
	}
	t, _ := template.New("page").Parse(`
{{.SiteBegin}}
{{.Header}}
<p>Welcome to the Northwind site. {{.Version}}</p>
{{.SiteEnd}}
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
		fmt.Fprint(&sb, "<ul>")
		for _, emp := range employees {
			fmt.Fprintf(&sb, "<li><a href=\"/employees/%d\">%s</a></li>", emp.Id, emp.FirstName+" "+emp.LastName)
		}
		fmt.Fprint(&sb, "</ul>")
		fmt.Fprint(w, SiteBegin() + Header()+sb.String() + SiteEnd())
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
		fmt.Fprint(w, SiteBegin() + Header()+sb.String() + SiteEnd())
	}
}
