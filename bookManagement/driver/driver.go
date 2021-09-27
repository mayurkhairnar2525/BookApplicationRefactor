package driver

import (
	"database/sql"
	"fmt"
)

const (
	DriverName = "mysql"
	ROOT       = "root"
	PASSWORD   = "12345678"
	HOST       = "0.0.0.0"
	PORT       = "9090"
	DBNAME     = "bookstore"
)

var db *sql.DB


func ConnectDB () *sql.DB{
	db, err := sql.Open(DriverName, ROOT+":"+PASSWORD+"@tcp("+HOST+":"+PORT+")/"+DBNAME)
	if err != nil {
		fmt.Println("error occurred", err)
	}
	fmt.Println("DB Connected")
	return db
}
