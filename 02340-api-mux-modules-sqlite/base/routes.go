package base

import (
	"fmt"
	"net/http"
)

func (app *App) handleHomeRoute(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>Northwind API</h1>")
}