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
	<li>get all books: <a href="http://localhost:9003/api/books">http://localhost:9003/api/books</a></li>
</ul>
	`)
}

func (app *App) handleGetBooksRoute(w http.ResponseWriter, _ *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	books := []string{"Homo Deux", "Sapiens"}
	json.NewEncoder(w).Encode(books)
}

func (app *App) handleGetSingleBookRoute(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r) 
	id := params["id"]
	w.Header().Set("Content-Type", "application/json")
	books := []string{"book " + string(id)}
	json.NewEncoder(w).Encode(books)
}