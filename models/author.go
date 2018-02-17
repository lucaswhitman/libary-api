package models

import (
	"errors"
)

type Author struct {
	Id        int
	FirstName string
	LastName  string
}

func ValidateAuthor(a Author, isNew bool) error {
	if a.Id <= 0 && isNew == false {
		return errors.New("Id must be greater than 0")
	}
	if a.Id != 0 && isNew == true {
		return errors.New("Id cannot be set on new author")
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
