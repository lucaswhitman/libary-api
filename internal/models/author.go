package main

import (
	"errors"
)

type Author struct {
	ID        int
	FirstName string
	LastName  string
}

func ValidateAuthor(a Author, isNew bool) error {
	if a.ID <= 0 && isNew == false {
		return errors.New("ID must be greater than 0")
	}
	if a.ID != 0 && isNew == true {
		return errors.New("ID cannot be set on new author")
	}
	if a.FirstName == "" && a.LastName == "" {
		return errors.New("Firstname and Lastname cannot be empty")
	}
	if a.FirstName == "" {
		return errors.New("FirstName cannot be empty")
	}
	if a.LastName == "" {
		return errors.New("LastName cannot be empty")
	}
	return nil
}
