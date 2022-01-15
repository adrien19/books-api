package api

import (
	model "books-api/model"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strconv"
)

func BooksIndex(w http.ResponseWriter, r *http.Request) {
	jsonBook, err := json.Marshal(model.StoreBooks)
	if err != nil {
		fmt.Println("failed - unable to process the request")
		w.WriteHeader(500)
		return
	}

	w.Write([]byte(jsonBook))
}

func BooksCreate(w http.ResponseWriter, r *http.Request) {
	var book model.Book
	err := json.NewDecoder(r.Body).Decode(&book)
	if err != nil {
		fmt.Println("failed to marshal")
		w.WriteHeader(400)
		return
	}
	book.ID = strconv.Itoa(len(model.StoreBooks) + 1)
	model.StoreBooks = append(model.StoreBooks, book)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(201)
	if err := json.NewEncoder(w).Encode(book); err != nil {
		fmt.Println("failed to encode book")
	}
}

func Book(w http.ResponseWriter, r *http.Request) {
	var bookID map[string]string
	var resp model.Book
	err := json.NewDecoder(r.Body).Decode(&bookID)
	if err != nil {
		fmt.Println("failed to marshal book id")
		w.WriteHeader(400)
		return
	}
	bookIndex, err := findBook(bookID["id"])
	if err != nil {
		fmt.Println("failed - book", err)
		w.WriteHeader(400)
		return
	}
	if bookIndex >= 0 {
		resp = model.StoreBooks[bookIndex]
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	if err := json.NewEncoder(w).Encode(resp); err != nil {
		fmt.Println("failed to encode book")
	}
}

func BooksEdit(w http.ResponseWriter, r *http.Request) {
	var book model.Book
	err := json.NewDecoder(r.Body).Decode(&book)
	if err != nil {
		fmt.Println("failed to marshal book id")
		w.WriteHeader(400)
		return
	}
	bookIndex, err := findBook(book.ID)
	if err != nil {
		fmt.Println("failed - book", err)
		w.WriteHeader(400)
		return
	}
	if bookIndex >= 0 {
		model.StoreBooks[bookIndex] = book
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(201)
	if err := json.NewEncoder(w).Encode(book); err != nil {
		fmt.Println("failed to encode book")
	}
}

func BooksDeleteByID(w http.ResponseWriter, r *http.Request) {
	var bookID map[string]string
	var resp string
	err := json.NewDecoder(r.Body).Decode(&bookID)
	if err != nil {
		fmt.Println("failed to marshal book ind")
		w.WriteHeader(400)
		return
	}
	bookIndex, err := findBook(bookID["id"])
	if err != nil {
		fmt.Println("failed - book", err)
		w.WriteHeader(400)
		return
	}
	if bookIndex >= 0 {
		model.StoreBooks = append(model.StoreBooks[:bookIndex], model.StoreBooks[bookIndex+1:]...)
		resp = "\nBook " + bookID["id"] + " deleted\n"
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(204)
	if err := json.NewEncoder(w).Encode(resp); err != nil {
		fmt.Println("failed to encode book")
	}
}

func findBook(id string) (int, error) {
	// Finds the book with the given id
	for i := 0; i < len(model.StoreBooks); i++ {
		if id == model.StoreBooks[i].ID {
			return i, nil
		}
	}
	return -1, errors.New("not found") // not found
}
