package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValidUpdate(t *testing.T) {
	a := Author{
		FirstName: "Roger",
		LastName:  "Zelazny",
		ID:        1,
	}
	if err := ValidateAuthor(a, false); err != nil {
		t.Errorf("Unexpected error: %s", err.Error())
	}
}

func TestValidCreate(t *testing.T) {
	a := Author{
		FirstName: "Roger",
		LastName:  "Zelazny",
	}
	if err := ValidateAuthor(a, true); err != nil {
		t.Errorf("Unexpected error: %s", err.Error())
	}
}

func TestIDNotSetOnUpdateError(t *testing.T) {
	a := Author{
		FirstName: "Roger",
		LastName:  "Zelazny",
	}
	expectedMessage := "ID must be greater than 0"
	if err := ValidateAuthor(a, false); err != nil {
		assert.Equal(t, expectedMessage, err.Error(), "Unexpected error message.")
	} else {
		t.Errorf("Expected: %s", expectedMessage)
	}
}

func TestNoIDOnCreateError(t *testing.T) {
	a := Author{
		FirstName: "Roger",
		LastName:  "Zelazny",
		ID:        1,
	}
	expectedMessage := "ID cannot be set on new author"
	if err := ValidateAuthor(a, true); err != nil {
		assert.Equal(t, expectedMessage, err.Error(), "Unexpected error message.")
	} else {
		t.Errorf("Expected: %s", expectedMessage)
	}
}

func TestEmptyNames(t *testing.T) {
	a := Author{
		FirstName: "",
		LastName:  "",
	}
	expectedMessage := "Firstname and Lastname cannot be empty"
	if err := ValidateAuthor(a, true); err != nil {
		assert.Equal(t, expectedMessage, err.Error(), "Unexpected error message.")
	} else {
		t.Errorf("Expected: %s", expectedMessage)
	}
}

func TestEmptyFirstName(t *testing.T) {
	a := Author{
		FirstName: "",
		LastName:  "Zelazny",
	}
	expectedMessage := "FirstName cannot be empty"
	if err := ValidateAuthor(a, true); err != nil {
		assert.Equal(t, expectedMessage, err.Error(), "Unexpected error message.")
	} else {
		t.Errorf("Expected: %s", expectedMessage)
	}
}

func TestEmptyLastName(t *testing.T) {
	a := Author{
		FirstName: "Roger",
		LastName:  "",
	}
	expectedMessage := "LastName cannot be empty"
	if err := ValidateAuthor(a, true); err != nil {
		assert.Equal(t, expectedMessage, err.Error(), "Unexpected error message.")
	} else {
		t.Errorf("Expected: %s", expectedMessage)
	}
}
