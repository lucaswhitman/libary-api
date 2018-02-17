// model.go

package main

import (
	"database/sql"
	"errors"
)

func (b *Book) getBook(db *sql.DB) error {
	return errors.New("Not implemented")
}

func (b *Book) updateBook(db *sql.DB) error {
	return errors.New("Not implemented")
}

func (b *Book) deleteBook(db *sql.DB) error {
	return errors.New("Not implemented")
}

func (b *Book) createBook(db *sql.DB) error {
	return errors.New("Not implemented")
}

func getBooks(db *sql.DB, start, count int) ([]Book, error) {
	return nil, errors.New("Not implemented")
}
