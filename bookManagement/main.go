package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"github.com/mayurkhairnar2525/bookManagement/controllers"
	"github.com/mayurkhairnar2525/bookManagement/driver"
	"github.com/mayurkhairnar2525/bookManagement/models"

	"net/http"
)

var books []models.BookManagement


var db *sql.DB
var err error


func main() {
	db = driver.ConnectDB()
	controllers := controllers.Controllers{}


	router := mux.NewRouter()
	router.HandleFunc("/books", controllers.GetBooks(db)).Methods("GET")
	router.HandleFunc("/books",controllers.CreateBook(db)).Methods("POST")
	router.HandleFunc("/books/{author}", controllers.GetBook(db)).Methods("GET")
	router.HandleFunc("/books/", controllers.UpdateBook(db)).Methods("PUT")
	router.HandleFunc("/books/{id}", controllers.DeleteBook(db)).Methods("DELETE")
	fmt.Println("Server is on port 8090:")
	http.ListenAndServe(":8090", router)
}
