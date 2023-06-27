package main

import (
	"encoding/json"
	"math/rand"
	"strconv"

	"net/http"

	"log"

	"github.com/gorilla/mux"
)

type Book struct {
	ID     string  `json:"id"`
	Isbn   string  `json:"isbn"`
	Title  string  `json:"title"`
	Author *Author `json:"author"`
}
type Author struct {
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
}

var books []Book

//function to get all books
func getBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(books)
}

func getBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for _, Item := range books {
		if Item.ID == params["id"] {
			json.NewEncoder(w).Encode(Item)
			return
		}
	}
	json.NewEncoder(w).Encode(&Book{})

}

func createBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var nayabook Book
	_ = json.NewDecoder(r.Body).Decode(&nayabook)
	nayabook.ID = strconv.Itoa(rand.Intn(10000000))
	books = append(books, nayabook)
	json.NewEncoder(w).Encode(nayabook)
}

func updateBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range books {
		if item.ID == params["id"] {
			books = append(books[:index], books[index+1:]...)
			var book Book
			_ = json.NewDecoder(r.Body).Decode(&book)
			book.ID = params["id"]
			books = append(books, book)
			json.NewEncoder(w).Encode(book)
			return
		}

	}

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
	//starting the router
	r := mux.NewRouter()
	books = append(books, Book{ID: "1", Isbn: "12226", Title: "How to make Crud", Author: &Author{FirstName: "Santosh", LastName: "Mukherjee"}})
	books = append(books, Book{ID: "2", Isbn: "45566", Title: "Go Docker", Author: &Author{FirstName: "Himanshu", LastName: "Ranjan"}})
	books = append(books, Book{ID: "3", Isbn: "2524", Title: "Go For Begineers", Author: &Author{FirstName: "Prashant", LastName: "Singh"}})

	//Creating Handlers it will establish endpoints between our API's
	r.HandleFunc("/books", getBooks).Methods("GET")
	r.HandleFunc("/books/{id}", getBook).Methods("GET")
	r.HandleFunc("/books", createBook).Methods("POST")
	r.HandleFunc("/books/{id}", updateBook).Methods("PUT")
	r.HandleFunc("/books/{id}", deleteBook).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8004", r))
}
