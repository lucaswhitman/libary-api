package models

import (
	"database/sql"
)

func (a *Author) getAuthor(db *sql.DB) error {
	return db.QueryRow("SELECT first_name, last_name FROM authors WHERE id=$1",
		a.Id).Scan(&a.FirstName, &a.LastName)
}

func (a *Author) updateAuthor(db *sql.DB) error {
	_, err :=
		db.Exec("UPDATE authors SET first_name=$1, last_name=$2 WHERE id=$3",
			a.FirstName, a.LastName, a.Id)

	return err
}

func (a *Author) deleteAuthor(db *sql.DB) error {
	_, err := db.Exec("DELETE FROM authors WHERE id=$1", a.Id)

	return err
}

func (a *Author) createAuthor(db *sql.DB) error {
	err := db.QueryRow(
		"INSERT INTO authors(first_name, last_name) VALUES($1, $2) RETURNING id",
		a.FirstName, a.LastName).Scan(&a.Id)

	if err != nil {
		return err
	}

	return nil
}

func getAuthors(db *sql.DB, start, count int) ([]Author, error) {
	rows, err := db.Query(
		"SELECT id, first_name, last_name FROM authors LIMIT $1 OFFSET $2",
		count, start)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	Authors := []Author{}

	for rows.Next() {
		var a Author
		if err := rows.Scan(&a.Id, &a.FirstName, &a.LastName); err != nil {
			return nil, err
		}
		Authors = append(Authors, a)
	}

	return Authors, nil
}
