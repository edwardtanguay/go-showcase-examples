package base

import (
	"database/sql"
	"log"
	"fmt"
	"net/http"

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
	http.HandleFunc("/", app.handleHomeRoute)
	http.HandleFunc("/employees", app.handleEmployeesRoute)
	fmt.Printf("API running at http://localhost:%d\n", app.Port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", app.Port), nil))
}
