package main

import (
	"fmt"
	"net/http"
)

func main() {
	port := 4392

	http.HandleFunc("/get-excel-file", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			http.Error(w, "could not find resource", http.StatusNotFound)
			return
		} else {
			html := `
			This is <b>html</b>.	
			`
			w.Header().Set("Content-Type", "text/html")
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(html))
		}

	})

	fmt.Printf("listening at http://localhost:%d", port)
	http.ListenAndServe(fmt.Sprintf(":%d", port), nil)

}
