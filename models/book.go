package models

import (
	"fmt"
	"database/sql"
	"encoding/json"
	"errors"
	"strings"
	"time"
)

const PUBLISH_DATE_FORMAT = "2006-01-02"

type JsonPublishDate time.Time

type StatusEnum string

const (
	CheckedIn  StatusEnum = "CheckedIn"
	CheckedOut StatusEnum = "CheckedOut"
)

type Book struct {
	ID          int             `json:"id,omitempty"`
	Title       string          `json:"title"`
	Author      string          `json:"author"`
	Publisher   string          `json:"publisher"`
	PublishDate JsonPublishDate `json:"publishDate"`
	Rating      int             `json:"rating"`
	Status      StatusEnum      `json:"status"`
}

func (b *Book) Validate(isNew bool) error {
	if b.ID <= 0 && isNew == false {
		return errors.New("ID must be greater than 0")
	}
	if b.ID != 0 && isNew == true {
		return errors.New("ID cannot be set on new book")
	}
	if b.Title == "" {
		return errors.New("Title cannot be empty")
	}
	if b.Author == "" {
		return errors.New("Author cannot be empty")
	}
	if b.Publisher == "" {
		return errors.New("Publisher cannot be empty")
	}
	if time.Time(b.PublishDate).After(time.Now()) {
		return errors.New("Cannot add unpublished books")
	}
	if b.Rating < 1 {
		return errors.New("Rating must be in range 1-3")
	}
	if b.Rating > 3 {
		return errors.New("Rating must be in range 1-3")
	}
	if b.Status != CheckedIn && b.Status != CheckedOut {
		return errors.New("Invalid status, valid statuses: CheckedIn, CheckedOut")
	}
	return nil
}

func (b *Book) GetBook(db *sql.DB) error {
	return db.QueryRow("SELECT title, author, publisher, publish_date, rating, status FROM books WHERE id=$1",
		b.ID).Scan(&b.Title, &b.Author, &b.Publisher, &b.PublishDate, &b.Rating, &b.Status)
}

func (b *Book) UpdateBook(db *sql.DB) error {
	_, err :=
		db.Exec("UPDATE books SET title=$1, author=$2, publisher=$3, publish_date=$4, rating=$5, status=$6 WHERE id=$7",
			b.Title, b.Author, b.Publisher, time.Time(b.PublishDate), b.Rating, b.Status, b.ID)

	return err
}

func (b *Book) DeleteBook(db *sql.DB) error {
	_, err := db.Exec("DELETE FROM books WHERE id=$1", b.ID)

	return err
}

func (b *Book) CreateBook(db *sql.DB) error {
	err := db.QueryRow(
		"INSERT INTO books(title, author, publisher, publish_date, rating, status) VALUES($1, $2, $3, $4, $5, $6) RETURNING id",
		b.Title, b.Author, b.Publisher, time.Time(b.PublishDate), b.Rating, b.Status).Scan(&b.ID)

	if err != nil {
		return err
	}

	return nil
}

func GetBooks(db *sql.DB, start, count int) ([]Book, error) {
	rows, err := db.Query(
		"SELECT id, title, author, publisher, publish_date, rating, status FROM books LIMIT $1 OFFSET $2",
		count, start)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	Books := []Book{}

	for rows.Next() {
		var b Book
		if err := rows.Scan(&b.ID, &b.Title, &b.Author, &b.Publisher, &b.PublishDate, &b.Rating, &b.Status); err != nil {
			return nil, err
		}
		Books = append(Books, b)
	}

	return Books, nil
}

// imeplement Marshaler und Unmarshaler interface
func (j *JsonPublishDate) UnmarshalJSON(b []byte) error {
	s := strings.Trim(string(b), "\"")
	t, err := time.Parse("2006-01-02", s)
	if err != nil {
		return err
	}
	*j = JsonPublishDate(t)
	return nil
}

func (j JsonPublishDate) MarshalJSON() ([]byte, error) {
	return json.Marshal(j)
}