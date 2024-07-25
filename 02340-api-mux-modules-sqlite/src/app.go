package src

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
	Router *mux.Router
}

func (app *App) Initialize() error {
	var err error
	app.DB, err = sql.Open("sqlite3", "./data/northwind.sqlite")
	if err != nil {
		return fmt.Errorf("error opening database: %v", err)
	}
	app.Router = mux.NewRouter()
	return nil
}

func (app *App) initializeRoutes() {
	app.Router.HandleFunc("/", app.handleHomeRoute).Methods("GET")
	app.Router.HandleFunc("/employees", app.handleGetEmployeesRoute).Methods("GET")
	app.Router.HandleFunc("/employees/{id}", app.handleGetSingleEmployeeRoute).Methods("GET")
}

func (app *App) Run() {
	app.initializeRoutes()

	staticDir := http.Dir("./static")
	staticFileHandler := http.StripPrefix("/", http.FileServer(staticDir))
	app.Router.PathPrefix("/").Handler(staticFileHandler).Methods("GET")

	http.Handle("/", app.Router)
	fmt.Printf("site running at http://localhost:%d\n", app.Port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", app.Port), app.Router))
}