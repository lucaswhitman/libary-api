// model.go

package main

import (
	"database/sql"
	"errors"
)

func (a *Author) getAuthor(db *sql.DB) error {
	return errors.New("Not implemented")
}

func (a *Author) updateAuthor(db *sql.DB) error {
	return errors.New("Not implemented")
}

func (a *Author) deleteAuthor(db *sql.DB) error {
	return errors.New("Not implemented")
}

func (a *Author) createAuthor(db *sql.DB) error {
	return errors.New("Not implemented")
}

func getAuthors(db *sql.DB, start, count int) ([]Author, error) {
	return nil, errors.New("Not implemented")
}
