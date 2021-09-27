package book

import (
	"database/sql"
	"github.com/mayurkhairnar2525/bookManagement/models"
	"log"
)

type BookRepository struct{}

func (b BookRepository) GetBooks(db *sql.DB, book models.BookManagement, books []models.BookManagement) ([]models.BookManagement, error) {
	result, err := db.Query("SELECT id, name, author, prices,available,lauchedyear, isbn,stock from bookmanagement")
	if err != nil {
		log.Println("err", err)
	}
	for result.Next() {
		err = result.Scan(&book.ID, &book.Name, &book.Author, &book.Prices, &book.Available, &book.LaunchedYear, &book.Isbn, &book.Stock)
		books = append(books, book)
	}
	if err != nil {
		log.Println("err:", err)
	}
	return books, nil
}

func (b BookRepository) GetBook(db *sql.DB, book models.BookManagement, id int) (models.BookManagement, error) {

	rows := db.QueryRow("SELECT id, name, author, prices, available, pagequality, lauchedyear, isbn,stock FROM bookmanagement WHERE author = ?")
	err := rows.Scan(&book.ID, &book.Name, &book.Author, &book.Prices, &book.Available, &book.PageQuality, &book.LaunchedYear, &book.Isbn, &book.Stock)

	return book, err
}

func (b BookRepository) CreateBook(db *sql.DB, book models.BookManagement) (int, error) {
	err := db.QueryRow("INSERT INTO bookmanagement(id,name,author,prices,available,pagequality,lauchedyear,isbn,stock) VALUES(?,?,?,?,?,?,?,?,?)", book.ID, book.Name, book.Author, book.Prices, book.Available, book.PageQuality, book.LaunchedYear, book.Isbn, book.Stock)
	if err == nil {
		log.Println("err occurred", err)
	}
	return book.ID, nil
}

func (b BookRepository) UpdateBook(db *sql.DB, book models.BookManagement) (int64, error) {
	result, err := db.Exec("UPDATE bookmanagement SET name =?, author=?,prices=?,available=?,pagequality=?,lauchedyear=?,stock=? WHERE isbn = ?",
		 &book.Name, &book.Author, &book.Prices, &book.Available, &book.PageQuality, &book.LaunchedYear, &book.Isbn, &book.Stock)
	if err != nil {
		log.Println("err", err)
	}
	rowsUpdated, err := result.RowsAffected()
	if err != nil {
		log.Println("err", err)
	}
	return rowsUpdated, nil
}

func (b BookRepository) DeleteBook(db *sql.DB, id int) (int64, error) {
	result, err := db.Exec("DELETE FROM bookmanagement WHERE id = ?",id)
	if err != nil {
		return 0, err
	}

	rowsDeleted, err := result.RowsAffected()
	if err != nil {
		return 0, err
	}
	return rowsDeleted, nil
}

