package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/yaonkey/gobookstore/pkg/models"
	"github.com/yaonkey/gobookstore/pkg/utils"
)

var NewBook models.Book

func CreateBook(w http.ResponseWriter, r *http.Request) {
	createBook := &models.Book{}
	utils.ParseBody(r, createBook)
	b := createBook.CreateBook()
	res, err := json.Marshal(b)
	if err != nil {
		sendError(w, err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Created book: %v", res)
}

func GetBook(w http.ResponseWriter, r *http.Request) {
	newBooks := models.GetAllBooks()
	res, err := json.Marshal(newBooks)
	if err != nil {
		sendError(w, err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Books: %v", res)
}

func GetBookById(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	bookId := params["bookId"]
	id, err := strconv.ParseUint(bookId, 0, 0)
	if err != nil {
		sendError(w, err)
	}
	bookDetails, _ := models.GetBookById(id)
	res, err := json.Marshal(bookDetails)
	if err != nil {
		sendError(w, err)
	}
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w, "Book by id %v: %v", id, res)

}

func DeleteBook(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	bookId := params["bookId"]
	id, err := strconv.ParseUint(bookId, 0, 0)
	if err != nil {
		sendError(w, err)
	}
	book := models.DeleteBook(id)
	res, err := json.Marshal(book)
	if err != nil {
		sendError(w, err)
	}
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w, "Book by id %v: %v", id, res)
}

func UpdateBook(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var updatedBook = &models.Book{}
	bookId := params["bookId"]
	utils.ParseBody(r, updatedBook)

	id, err := strconv.ParseUint(bookId, 0, 0)
	if err != nil {
		sendError(w, err)
	}

	bookDetails, db := models.GetBookById(id)
	if updatedBook.Name != "" {
		bookDetails.Name = updatedBook.Name
	}
	if updatedBook.Author != "" {
		bookDetails.Author = updatedBook.Author
	}
	if updatedBook.Publication != "" {
		bookDetails.Publication = updatedBook.Publication
	}

	db.Save(&bookDetails)
	res, err := json.Marshal(bookDetails)
	if err != nil {
		sendError(w, err)
	}
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w, "Book by id %v: %v", id, res)
}

func sendError(w http.ResponseWriter, err error) {
	w.WriteHeader(http.StatusInternalServerError)
	fmt.Fprintf(w, "Internal error: %v", err)
}
