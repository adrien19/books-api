package main

import (
	api "books-api/api"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {

	m := mux.NewRouter()

	m.HandleFunc("/books", api.BooksIndex).Methods("GET")
	m.HandleFunc("/book", api.Book).Methods("GET")
	m.HandleFunc("/books/create", api.BooksCreate).Methods("POST")
	m.HandleFunc("/books/edit", api.BooksEdit).Methods("PUT")
	m.HandleFunc("/books/delete", api.BooksDeleteByID).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":3000", m))
}
