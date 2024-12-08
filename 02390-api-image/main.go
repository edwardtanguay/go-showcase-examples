package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/logo", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			fmt.Println("ok")
		} else {
			http.Error(w, "resource not found", http.StatusNotFound)
			return
		}
	})

	port := 4489
	fmt.Printf("Listening at http://localhost:%d\n", port)
	http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
}
