package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

type User struct {
	Login    string `json:"login"`
	Password string `json:"password`
}

func loginHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain")
	w.Write([]byte("Welcome!"))
}

func authentication(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var newUser User

		err := json.NewDecoder(r.Body).Decode(&newUser)
		if err != nil {
			fmt.Println(err)
		}

		if newUser.Login != "hans" || newUser.Password != "123" {
			http.Error(w, "no access", http.StatusUnauthorized)
			return
		}
		next.ServeHTTP(w, r)
	})
}

func main() {
	port := 7283
	mux := http.NewServeMux()
	middleware := authentication(http.HandlerFunc(loginHandler))
	mux.Handle("/login", middleware)
	fmt.Printf("listening at: http://localhost:%d", port)
	http.ListenAndServe(":"+strconv.Itoa(port), mux)
}
