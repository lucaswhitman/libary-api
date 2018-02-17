package main

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

const TEST_DATE_STRING = "2014-11-12T11:45:26.371Z"

func TestBookValidUpdate(t *testing.T) {
	d, _ := time.Parse(PUBLISH_DATE_FORMAT, TEST_DATE_STRING)
	b := Book{
		ID:          1,
		Title:       "The Great Book Of Amber",
		AuthorID:    1,
		Publisher:   "Harper Voyager",
		PublishDate: d,
		Rating:      1,
		Status:      CheckedIn,
	}
	if err := ValidateBook(b, false); err != nil {
		t.Errorf("Unexpected error: %s", err.Error())
	}
}

func TestBookValidCreate(t *testing.T) {
	d, _ := time.Parse(PUBLISH_DATE_FORMAT, TEST_DATE_STRING)
	b := Book{
		Title:       "The Great Book Of Amber",
		AuthorID:    1,
		Publisher:   "Harper Voyager",
		PublishDate: d,
		Rating:      1,
		Status:      CheckedIn,
	}
	if err := ValidateBook(b, true); err != nil {
		t.Errorf("Unexpected error: %s", err.Error())
	}
}

func TestBookIDNotSetOnCreate(t *testing.T) {
	d, _ := time.Parse(PUBLISH_DATE_FORMAT, TEST_DATE_STRING)
	b := Book{
		ID:          1,
		Title:       "The Great Book Of Amber",
		AuthorID:    1,
		Publisher:   "Harper Voyager",
		PublishDate: d,
		Rating:      1,
		Status:      CheckedIn,
	}
	expectedMessage := "ID cannot be set on new book"
	if err := ValidateBook(b, true); err != nil {
		assert.Equal(t, expectedMessage, err.Error(), "Unexpected error message.")
	} else {
		t.Errorf("Expected: %s", expectedMessage)
	}
}

func TestBookIDSetOnUpdate(t *testing.T) {
	d, _ := time.Parse(PUBLISH_DATE_FORMAT, TEST_DATE_STRING)
	b := Book{
		Title:       "The Great Book Of Amber",
		AuthorID:    1,
		Publisher:   "Harper Voyager",
		PublishDate: d,
		Rating:      1,
		Status:      CheckedIn,
	}
	expectedMessage := "ID must be greater than 0"
	if err := ValidateBook(b, false); err != nil {
		assert.Equal(t, expectedMessage, err.Error(), "Unexpected error message.")
	} else {
		t.Errorf("Expected: %s", expectedMessage)
	}
}

func TestBookTitleNotEmpty(t *testing.T) {
	d, _ := time.Parse(PUBLISH_DATE_FORMAT, TEST_DATE_STRING)
	b := Book{
		Title:       "",
		AuthorID:    1,
		Publisher:   "Harper Voyager",
		PublishDate: d,
		Rating:      1,
		Status:      CheckedIn,
	}
	expectedMessage := "Title cannot be empty"
	if err := ValidateBook(b, true); err != nil {
		assert.Equal(t, expectedMessage, err.Error(), "Unexpected error message.")
	} else {
		t.Errorf("Expected: %s", expectedMessage)
	}
}

func TestBookAuthorIDNotZero(t *testing.T) {
	d, _ := time.Parse(PUBLISH_DATE_FORMAT, TEST_DATE_STRING)
	b := Book{
		Title:       "The Great Book Of Amber",
		AuthorID:    0,
		Publisher:   "Harper Voyager",
		PublishDate: d,
		Rating:      1,
		Status:      CheckedIn,
	}
	expectedMessage := "AuthorID must be greater than 0"
	if err := ValidateBook(b, true); err != nil {
		assert.Equal(t, expectedMessage, err.Error(), "Unexpected error message.")
	} else {
		t.Errorf("Expected: %s", expectedMessage)
	}
}

func TestBookPublisherNotEmpty(t *testing.T) {
	d, _ := time.Parse(PUBLISH_DATE_FORMAT, TEST_DATE_STRING)
	b := Book{
		Title:       "The Great Book Of Amber",
		AuthorID:    1,
		Publisher:   "",
		PublishDate: d,
		Rating:      1,
		Status:      CheckedIn,
	}
	expectedMessage := "Publisher cannot be empty"
	if err := ValidateBook(b, true); err != nil {
		assert.Equal(t, expectedMessage, err.Error(), "Unexpected error message.")
	} else {
		t.Errorf("Expected: %s", expectedMessage)
	}
}

func TestBookNotPublishedYet(t *testing.T) {
	now := time.Now()
	threeDays := time.Hour * 24 * 3
	futureDate := now.Add(threeDays)

	b := Book{
		Title:       "The Great Book Of Amber",
		AuthorID:    1,
		Publisher:   "Harper Voyager",
		PublishDate: futureDate,
		Rating:      1,
		Status:      CheckedIn,
	}
	expectedMessage := "Cannot add unpublished books"
	if err := ValidateBook(b, true); err != nil {
		assert.Equal(t, expectedMessage, err.Error(), "Unexpected error message.")
	} else {
		t.Errorf("Expected: %s", expectedMessage)
	}
}

func TestBookRatingTooLow(t *testing.T) {
	d, _ := time.Parse(PUBLISH_DATE_FORMAT, TEST_DATE_STRING)
	b := Book{
		Title:       "The Great Book Of Amber",
		AuthorID:    1,
		Publisher:   "Harper Voyager",
		PublishDate: d,
		Rating:      0,
		Status:      CheckedIn,
	}
	expectedMessage := "Rating must be in range 1-3"
	if err := ValidateBook(b, true); err != nil {
		assert.Equal(t, expectedMessage, err.Error(), "Unexpected error message.")
	} else {
		t.Errorf("Expected: %s", expectedMessage)
	}
}

func TestBookRatingTooHigh(t *testing.T) {
	d, _ := time.Parse(PUBLISH_DATE_FORMAT, TEST_DATE_STRING)
	b := Book{
		Title:       "The Great Book Of Amber",
		AuthorID:    1,
		Publisher:   "Harper Voyager",
		PublishDate: d,
		Rating:      110,
		Status:      CheckedIn,
	}
	expectedMessage := "Rating must be in range 1-3"
	if err := ValidateBook(b, true); err != nil {
		assert.Equal(t, expectedMessage, err.Error(), "Unexpected error message.")
	} else {
		t.Errorf("Expected: %s", expectedMessage)
	}
}

func TestBookStatusEmpty(t *testing.T) {
	d, _ := time.Parse(PUBLISH_DATE_FORMAT, TEST_DATE_STRING)
	b := Book{
		Title:       "The Great Book Of Amber",
		AuthorID:    1,
		Publisher:   "Harper Voyager",
		PublishDate: d,
		Rating:      1,
		Status:      "",
	}
	expectedMessage := "Invalid status, valid statuses: CheckedIn, CheckedOut"
	if err := ValidateBook(b, true); err != nil {
		assert.Equal(t, expectedMessage, err.Error(), "Unexpected error message.")
	} else {
		t.Errorf("Expected: %s", expectedMessage)
	}
}

func TestBookStatusStringConversion(t *testing.T) {
	d, _ := time.Parse(PUBLISH_DATE_FORMAT, TEST_DATE_STRING)
	b := Book{
		Title:       "The Great Book Of Amber",
		AuthorID:    1,
		Publisher:   "Harper Voyager",
		PublishDate: d,
		Rating:      1,
		Status:      "CheckedIn",
	}
	if err := ValidateBook(b, true); err != nil {
		t.Errorf("Unexpected error: %s", err.Error())
	}
}
