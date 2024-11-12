package db

import (
	"database/sql"
	"fmt"

	_ "github.com/microsoft/go-mssqldb"
)

// Storage represents the database storage.
type Storage struct {
	DB *sql.DB // changed line
}

// NewStorage creates a new Storage instance and connects to the database.
func NewStorage(connString string) (*Storage, error) {
	db, err := sql.Open("sqlserver", connString)
	if err != nil {
		return nil, err
	}
	// defer db.Close()

	err = db.Ping()
	if err != nil {
		return nil, err
	}
	fmt.Println("DB Connected!")

	return &Storage{DB: db}, nil // changed line
}
