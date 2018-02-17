package main

import (
	"errors"
	"time"
)

const PUBLISH_DATE_FORMAT = "2006-01-02"

type StatusEnum string

const (
	CheckedIn  StatusEnum = "CheckedIn"
	CheckedOut StatusEnum = "CheckedOut"
)

type Book struct {
	ID          int
	Title       string
	AuthorID    int
	Publisher   string
	PublishDate time.Time
	Rating      int
	Status      StatusEnum
}

func ValidateBook(b Book, isNew bool) error {
	if b.ID <= 0 && isNew == false {
		return errors.New("ID must be greater than 0")
	}
	if b.ID != 0 && isNew == true {
		return errors.New("ID cannot be set on new book")
	}
	if b.Title == "" {
		return errors.New("Title cannot be empty")
	}
	if b.AuthorID <= 0 {
		return errors.New("AuthorID must be greater than 0")
	}
	if b.Publisher == "" {
		return errors.New("Publisher cannot be empty")
	}
	if b.PublishDate.After(time.Now()) {
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
