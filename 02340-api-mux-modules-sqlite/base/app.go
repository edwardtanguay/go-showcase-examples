package base

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/mattn/go-sqlite3"
)

type App struct {
	DB   *sql.DB
	Port int
}

func (app *App) Initialize() error {
	var err error
	app.DB, err = sql.Open("sqlite3", "./data/northwind.sqlite")
	if err != nil {
		return fmt.Errorf("error opening database: %v", err)
	}
	return nil
}

func (app *App) Run() {
	r := mux.NewRouter()
	r.HandleFunc("/", app.handleHomeRoute).Methods("GET")
	r.HandleFunc("/api/employees", app.handleGetEmployeesRoute).Methods("GET")
	r.HandleFunc("/api/employees/{id}", app.handleGetSingleEmployeeRoute).Methods("GET")
	fmt.Printf("API running at http://localhost:%d\n", app.Port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", app.Port), r))
}