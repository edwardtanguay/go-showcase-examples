package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Book struct {
	ID     string `json:"id"`
	Title  string `json:"title"`
	Author string `json:"author"`
	Year   string `json:"year"`
}

var books []Book

func getBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(books)
}

func getBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for _, item := range books {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&Book{})
}

func createBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var book Book
	_ = json.NewDecoder(r.Body).Decode(&book)
	books = append(books, book)
	json.NewEncoder(w).Encode(book)
}

func replaceBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range books {
		if item.ID == params["id"] {
			books = append(books[:index], books[index+1:]...)
			var book Book
			_ = json.NewDecoder(r.Body).Decode(&book)
			fmt.Printf("%#v\n", book)
			book.ID = params["id"]
			books = append(books, book)
			json.NewEncoder(w).Encode(book)
			return
		}
	}
	json.NewEncoder(w).Encode(books)
}

func editBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range books {
		if item.ID == params["id"] {
			var updates map[string]interface{}
			err := json.NewDecoder(r.Body).Decode(&updates)
			if err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}
			if title, ok := updates["title"].(string); ok {
				books[index].Title = title
			}
			if author, ok := updates["author"].(string); ok {
				books[index].Author = author
			}
			if year, ok := updates["year"].(string); ok {
				books[index].Year = year
			}

			json.NewEncoder(w).Encode(books[index])
			return
		}
	}

	http.Error(w, "Book not found", http.StatusNotFound)
}

func deleteBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range books {
		if item.ID == params["id"] {
			books = append(books[:index], books[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(books)
}

func main() {
	r := mux.NewRouter()

	// Mock data - @TODO: implement database
	books = append(books, Book{ID: "1", Title: "1984", Author: "George Orwell", Year: "1949"})
	books = append(books, Book{ID: "2", Title: "Brave New World", Author: "Aldous Huxley", Year: "1932"})

	// Route handlers / Endpoints
	r.HandleFunc("/api/books", getBooks).Methods("GET")
	r.HandleFunc("/api/books/{id}", getBook).Methods("GET")
	r.HandleFunc("/api/books", createBook).Methods("POST")
	r.HandleFunc("/api/books/{id}", replaceBook).Methods("PUT")
	r.HandleFunc("/api/books/{id}", editBook).Methods("PATCH")
	r.HandleFunc("/api/books/{id}", deleteBook).Methods("DELETE")

	fmt.Println("Listening at: http://localhost:9003/api/books")
	log.Fatal(http.ListenAndServe(":9003", r))
}
