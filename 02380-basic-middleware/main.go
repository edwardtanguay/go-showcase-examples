package main

import (
	"fmt"
	"net/http"
	"strconv"
)

type User struct {
	login    string
	password string
}

func loginHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain")
	w.Write([]byte("Welcome!"))
}

func main() {
	port := 7283
	http.HandleFunc("/login", loginHandler)
	fmt.Printf("listening at: http://localhost:%d", port)
	http.ListenAndServe(":"+strconv.Itoa(port), nil)
}
