package base

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
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
	r.HandleFunc("/api/books", app.handleGetBooksRoute).Methods("GET")
	r.HandleFunc("/api/books/{id}", app.handleGetSingleBookRoute).Methods("GET")
	fmt.Printf("API running at http://localhost:%d\n", app.Port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", app.Port), r))
}