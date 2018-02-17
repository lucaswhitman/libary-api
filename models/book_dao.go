package models

import "database/sql"

func (b *Book) getBook(db *sql.DB) error {
	return db.QueryRow("SELECT title, author_id, publisher, publish_date, rating, status FROM books WHERE id=$1",
		b.Id).Scan(&b.Title, &b.AuthorId, &b.Publisher, &b.PublishDate, &b.Rating, &b.Status)
}

func (b *Book) updateBook(db *sql.DB) error {
	_, err :=
		db.Exec("UPDATE books SET title=$1, author_id=$2, publisher=$3, publish_date=$4, rating=$5, status=$6 WHERE id=$7",
			b.Title, b.AuthorId, b.Publisher, b.PublishDate, b.Rating, b.Status, b.Id)

	return err
}

func (b *Book) deleteBook(db *sql.DB) error {
	_, err := db.Exec("DELETE FROM books WHERE id=$1", b.Id)

	return err
}

func (b *Book) createBook(db *sql.DB) error {
	err := db.QueryRow(
		"INSERT INTO books(title, author_id, publisher, publish_date, rating, status) VALUES($1, $2, $3, $4, $5, $6) RETURNING id",
		b.Title, b.AuthorId, b.Publisher, b.PublishDate, b.Rating, b.Status, b.Id).Scan(&b.Id)

	if err != nil {
		return err
	}

	return nil
}

func getBooks(db *sql.DB, start, count int) ([]Book, error) {
	rows, err := db.Query(
		"SELECT id, name, price FROM books LIMIT $1 OFFSET $2",
		count, start)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	Books := []Book{}

	for rows.Next() {
		var b book
		if err := rows.Scan(&b.ID, &b.Name, &b.Price); err != nil {
			return nil, err
		}
		Books = append(Books, b)
	}

	return Books, nil
}
