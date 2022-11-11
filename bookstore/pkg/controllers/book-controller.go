package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/ferseg/learning-go/bookstore/pkg/models"
	"github.com/ferseg/learning-go/bookstore/pkg/utils"
	_ "github.com/ferseg/learning-go/bookstore/pkg/utils"
	"github.com/gorilla/mux"
)

var TheBook models.Book

func GetAllBooks(w http.ResponseWriter, r *http.Request) {
	books := models.GetAllBooks()
	res, _ := json.Marshal(books)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func GetBookById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	bookId := vars["bookId"]
	id, err := strconv.ParseInt(bookId, 0, 0)
	if err != nil {
		fmt.Println("Error while parsing")
	}
	book, _ := models.GetBookById(id)
	res, _ := json.Marshal(book)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func CreateBook(w http.ResponseWriter, r *http.Request) {
    fmt.Println("It has reached this point")
	CreateBook := &models.Book{}
	utils.ParseBody(r, CreateBook )
	savedBook := CreateBook.CreateBook()
	res, _ := json.Marshal(savedBook)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func DeleteBook(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	bookId := vars["bookId"]
	id, err := strconv.ParseInt(bookId, 0, 0)
	if err != nil {
		fmt.Println("Error converting to int when deleting")
	}
	deletedBook := models.DeleteBook(id)

	res, _ := json.Marshal(deletedBook)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func UpdateBook(w http.ResponseWriter, r *http.Request) {
	updateBook := &models.Book{}
	utils.ParseBody(r, updateBook)
	vars := mux.Vars(r)
	id, err := strconv.ParseInt(vars["bookId"], 0, 0)
	if err != nil {
		fmt.Println("Error while parsing id")
	}
	existingBook, _ := models.GetBookById(id)
	if existingBook.ID > 0 {
		resultingBook := models.UpdateBook(id, updateBook)

		res, _ := json.Marshal(resultingBook)
		w.WriteHeader(http.StatusOK)
		w.Write(res)
	}
}
