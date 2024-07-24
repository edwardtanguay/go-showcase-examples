package base

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func (app *App) handleHomeRoute(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>Northwind API</h1>")
}

func (app *App) handleGetBooksRoute(w http.ResponseWriter, _ *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	books := []string{"Homo Deux", "Sapiens"}
	json.NewEncoder(w).Encode(books)
}