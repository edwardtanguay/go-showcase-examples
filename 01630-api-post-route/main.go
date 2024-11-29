package main

import (
	"fmt"
	"net/http"
	"strconv"
)

type Flashcard struct {
	Front string `json: "front"`
	Back  string `json: "back"`
}

func handlePostCreateFlashcard(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method, expected POST", http.StatusMethodNotAllowed)
		return
	}

}

func main() {
	port := 3988
	http.HandleFunc("/create-flashcard", handlePostCreateFlashcard)
	fmt.Printf("Server running at http://localhost:%s", strconv.Itoa(port))
	http.ListenAndServe(":" + strconv.Itoa(port), nil)
}