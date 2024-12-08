package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/logo", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			imagePathAndFileName := "images/logo.png"
			image, err := os.ReadFile(imagePathAndFileName)
			if err != nil {
				fmt.Println("bad image request")
				http.Error(w, "bad image", http.StatusBadRequest)
				return
			}
			w.WriteHeader(http.StatusOK)
			w.Header().Set("Content-Type", "application/octet-stream")
			w.Write(image)
		} else {
			http.Error(w, "resource not found", http.StatusNotFound)
			return
		}
	})

	port := 4489
	fmt.Printf("Listening at http://localhost:%d\n", port)
	http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
}
