package controllers

import (
	"database/sql"
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/mayurkhairnar2525/bookManagement/models"
	_ "github.com/mayurkhairnar2525/bookManagement/repository/book"
	book2 "github.com/mayurkhairnar2525/bookManagement/repository/book"
	"github.com/mayurkhairnar2525/bookManagement/utils"
	"log"
	"net/http"
	"strconv"
)

type Controllers struct{}

var books []models.BookManagement

func (c Controllers) GetBooks(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var book models.BookManagement
		var error models.Error

		books = []models.BookManagement{}
		bookRepo := book2.BookRepository{}
		books, err := bookRepo.GetBooks(db, book, books)
		if err != nil {
			error.Message = "server error"
			utils.SendError(w, http.StatusInternalServerError, error)
			return
		}
		w.Header().Set("Content-Type", "application/json")
	err = 	json.NewEncoder(w).Encode(books)
	if err!=nil{
		log.Println("err",err)
	}
	}
}

func (c Controllers) CreateBook(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var book models.BookManagement
		var bookID int

		var error models.Error

		err := json.NewDecoder(r.Body).Decode(&book)
		if err != nil {
			error.Message = "error occurred"
		}
		if book.Name == "" || book.Author == "" || book.Prices == 0 || book.Available == "" || book.PageQuality == "" || book.LaunchedYear == "" || book.Isbn == "" || book.Stock == 0 {
			error.Message = "Enter missing fields."
			utils.SendError(w, http.StatusBadRequest, error) //400
			return
		}
		bookRepo := book2.BookRepository{}
		bookID, err = bookRepo.CreateBook(db, book)
		if err != nil {
			error.Message = "Server error"
			utils.SendError(w, http.StatusInternalServerError, error) //500
			return
		}
		w.Header().Set("Content-Type", "text/plain")
		utils.SendSuccess(w, bookID)
	}
}

func (c Controllers) GetBook(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var book models.BookManagement
		var error models.Error

		params := mux.Vars(r)
		books = []models.BookManagement{}
		bookRepo := book2.BookRepository{}

		for author, booked := range books {
			if book.Author == params["author"] {
				err := json.NewEncoder(w).Encode(&booked)
				if err == nil {
					log.Println("err:", err)
				}
			}
			book, err := bookRepo.GetBook(db, booked, author)
			if err != nil {
				if err == sql.ErrNoRows {
					error.Message = "Not Found"
					utils.SendError(w, http.StatusNotFound, error)
					return
				} else {
					error.Message = "Server error"
					utils.SendError(w, http.StatusInternalServerError, error)
					return
				}
			}
			w.Header().Set("Content-Type", "application/json")
			utils.SendSuccess(w, book)
		}

	}
}

func (c Controllers) UpdateBook(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var book models.BookManagement
		var error models.Error

		err := json.NewDecoder(r.Body).Decode(&book)
		if err != nil {
			log.Println("err", err)
		}
		if book.ID == 0 || book.Name == "" || book.Author == "" || book.Prices == 0 || book.Available == "" || book.PageQuality == "" || book.LaunchedYear == "" || book.Isbn == "" || book.Stock == 0 {
			error.Message = "All fields are required."
			utils.SendError(w, http.StatusBadRequest, error)
			return
		}
		bookRepo := book2.BookRepository{}
		rowsUpdated, err := bookRepo.UpdateBook(db, book)
		if err != nil {
			error.Message = "Server error"
			utils.SendError(w, http.StatusInternalServerError, error) //500
			return
		}

		w.Header().Set("Content-Type", "text/plain")
		utils.SendSuccess(w, rowsUpdated)
	}
}

func (c Controllers) DeleteBook(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var error models.Error
		params := mux.Vars(r)
		bookRepo := book2.BookRepository{}
		id, err := strconv.Atoi(params["id"])
		if err!=nil{
			log.Println("err",err)
		}
		rowsDeleted, err := bookRepo.DeleteBook(db, id)
		if err != nil {
			error.Message = "Server error."
			utils.SendError(w, http.StatusInternalServerError, error) //500
			return
		}
		if rowsDeleted == 0 {
			error.Message = "Not Found"
			utils.SendError(w, http.StatusNotFound, error) //404
			return
		}
		w.Header().Set("Content-Type", "text/plain")
		utils.SendSuccess(w, rowsDeleted)
	}
}

