package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

type Flashcard struct {
	Id    string `json:"id"`
	Front string `json:"front"`
	Back  string `json:"back"`
}

func handlePostCreateFlashcard(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method, expected POST", http.StatusMethodNotAllowed)
		return
	}

	var flashcard Flashcard
	err := json.NewDecoder(r.Body).Decode(&flashcard)
	if err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
	}
	flashcard.Id = getSuuid()
	
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(flashcard)
}

func main() {
	port := 3988
	http.HandleFunc("/create-flashcard", handlePostCreateFlashcard)
	fmt.Printf("Server running at http://localhost:%s\n", strconv.Itoa(port))
	http.ListenAndServe(":"+strconv.Itoa(port), nil)
}
