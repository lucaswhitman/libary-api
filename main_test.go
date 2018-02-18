package main

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"strconv"
	"testing"

	"github.com/lucaswhitman/library-api/app"
)

func TestMain(m *testing.M) {
	a = app.App{}

	conf, err := getConf("./config_test.json")
	if err != nil {
		log.Fatal(err)
	}
	a.Initialize(conf.Database.Host, conf.Database.Port, conf.Database.Username, conf.Database.Password, conf.Database.DatabaseName)
	booksTableCreationQuery, err := ioutil.ReadFile("sql/createBooksTable.sql")
	if err != nil {
		log.Fatal(err)
	}
	if _, err := a.DB.Exec(string(booksTableCreationQuery)); err != nil {
		log.Fatal(err)
	}
	code := m.Run()

	clearTables()

	os.Exit(code)
}

func clearTables() {
	a.DB.Exec("DELETE FROM books")
	a.DB.Exec("ALTER SEQUENCE books_id_seq RESTART WITH 1")
}

func TestEmptyTable(t *testing.T) {
	clearTables()

	req, _ := http.NewRequest("GET", "/books", nil)
	response := executeRequest(req)

	checkResponseCode(t, http.StatusOK, response.Code)

	if body := response.Body.String(); body != "[]" {
		t.Errorf("Expected an empty array. Got %s", body)
	}
}

func executeRequest(req *http.Request) *httptest.ResponseRecorder {
	rr := httptest.NewRecorder()
	a.Router.ServeHTTP(rr, req)

	return rr
}

func checkResponseCode(t *testing.T, expected, actual int) {
	if expected != actual {
		t.Errorf("Expected response code %d. Got %d\n", expected, actual)
	}
}

func TestCreateBook(t *testing.T) {
	clearTables()

	payload := []byte(`{
		"title":"Hitchhikers Guide to the Galaxy",
		"author": "Douglas Adams",
		"publisher": "Pan Books",
		"publishDate": "1979-10-12T11:45:26.371Z",
		"rating": 3,
		"status": "CheckedOut"}`)

	req, _ := http.NewRequest("POST", "/books", bytes.NewBuffer(payload))
	response := executeRequest(req)

	checkResponseCode(t, http.StatusCreated, response.Code)

	var m map[string]interface{}
	json.Unmarshal(response.Body.Bytes(), &m)

	if m["title"] != "Hitchhikers Guide to the Galaxy" {
		t.Errorf("Expected product name to be 'Hitchhikers Guide to the Galaxy'. Got '%v'", m["name"])
	}
}

func TestCreateBookBadRequest(t *testing.T) {
	clearTables()

	payload := []byte(`{
		"author": "Douglas Adams",
		"publisher": "Pan Books",
		"publishDate": "1979-10-12T11:45:26.371Z",
		"rating": 3,
		"status": "CheckedOut"}`)

	req, _ := http.NewRequest("POST", "/books", bytes.NewBuffer(payload))
	response := executeRequest(req)

	checkResponseCode(t, http.StatusBadRequest, response.Code)

	var m map[string]string
	json.Unmarshal(response.Body.Bytes(), &m)
	if m["error"] != "Title cannot be empty" {
		t.Errorf("Expected the 'error' key of the response to be set to 'Title cannot be empty'. Got '%s'", m["error"])
	}
}

func TestGetNonExistentBook(t *testing.T) {
	clearTables()

	req, _ := http.NewRequest("GET", "/books/11", nil)
	response := executeRequest(req)

	checkResponseCode(t, http.StatusNotFound, response.Code)

	var m map[string]string
	json.Unmarshal(response.Body.Bytes(), &m)
	if m["error"] != "Book not found" {
		t.Errorf("Expected the 'error' key of the response to be set to 'Book not found'. Got '%s'", m["error"])
	}
}

func TestUpdateBook(t *testing.T) {
	clearTables()
	addBooks(1)

	req, _ := http.NewRequest("GET", "/books/1", nil)
	response := executeRequest(req)
	var originalBook map[string]interface{}
	json.Unmarshal(response.Body.Bytes(), &originalBook)

	payload := []byte(`{
		"id": 1,
		"title":"Hitchhikers Guide to the Galaxy",
		"author": "Douglas Adams",
		"publisher": "Pan Books",
		"publishDate": "1979-10-12T11:45:26.371Z",
		"rating": 3,
		"status": "CheckedOut"}`)

	req, _ = http.NewRequest("PUT", "/books/1", bytes.NewBuffer(payload))
	response = executeRequest(req)

	checkResponseCode(t, http.StatusOK, response.Code)

	var m map[string]interface{}
	json.Unmarshal(response.Body.Bytes(), &m)

	if m["id"] != originalBook["id"] {
		t.Errorf("Expected the id to remain the same (%v). Got %v", originalBook["id"], m["id"])
	}

	if m["title"] == originalBook["title"] {
		t.Errorf("Expected the title to change from '%v' to '%v'. Got '%v'", originalBook["title"], m["title"], m["title"])
	}

	if m["author"] == originalBook["author"] {
		t.Errorf("Expected the author to change from '%v' to '%v'. Got '%v'", originalBook["author"], m["author"], m["author"])
	}

	if m["publisher"] == originalBook["publisher"] {
		t.Errorf("Expected the publisher to change from '%v' to '%v'. Got '%v'", originalBook["publisher"], m["publisher"], m["publisher"])
	}

	if m["publishDate"] == originalBook["publishDate"] {
		t.Errorf("Expected the publishDate to change from '%v' to '%v'. Got '%v'", originalBook["publishDate"], m["publishDate"], m["publishDate"])
	}

	if m["rating"] == originalBook["rating"] {
		t.Errorf("Expected the rating to change from '%v' to '%v'. Got '%v'", originalBook["rating"], m["rating"], m["rating"])
	}

	if m["status"] == originalBook["status"] {
		t.Errorf("Expected the status to change from '%v' to '%v'. Got '%v'", originalBook["status"], m["status"], m["status"])
	}
}

func TestUpdateBookBadID(t *testing.T) {
	clearTables()
	addBooks(1)

	req, _ := http.NewRequest("GET", "/books/1", nil)
	response := executeRequest(req)
	var originalBook map[string]interface{}
	json.Unmarshal(response.Body.Bytes(), &originalBook)

	payload := []byte(`{
		"id": 20,
		"title":"Hitchhikers Guide to the Galaxy",
		"author": "Douglas Adams",
		"publisher": "Pan Books",
		"publishDate": "1979-10-12T11:45:26.371Z",
		"rating": 3,
		"status": "CheckedOut"}`)

	req, _ = http.NewRequest("PUT", "/books/1", bytes.NewBuffer(payload))
	response = executeRequest(req)

	checkResponseCode(t, http.StatusBadRequest, response.Code)

	var m map[string]string
	json.Unmarshal(response.Body.Bytes(), &m)
	if m["error"] != "Book ID in body does not match ID in URI" {
		t.Errorf("Expected the 'error' key of the response to be set to 'Book ID in body does not match ID in URI'. Got '%s'", m["error"])
	}
}

func TestUpdateBookBadRequest(t *testing.T) {
	clearTables()
	addBooks(1)

	req, _ := http.NewRequest("GET", "/books/1", nil)
	response := executeRequest(req)
	var originalBook map[string]interface{}
	json.Unmarshal(response.Body.Bytes(), &originalBook)

	payload := []byte(`{
		"id": 1,
		"title":"Hitchhikers Guide to the Galaxy",
		"publisher": "Pan Books",
		"publishDate": "1979-10-12T11:45:26.371Z",
		"rating": 3,
		"status": "CheckedOut"}`)

	req, _ = http.NewRequest("PUT", "/books/1", bytes.NewBuffer(payload))
	response = executeRequest(req)

	checkResponseCode(t, http.StatusBadRequest, response.Code)

	var m map[string]string
	json.Unmarshal(response.Body.Bytes(), &m)
	if m["error"] != "Author cannot be empty" {
		t.Errorf("Expected the 'error' key of the response to be set to 'Author cannot be empty'. Got '%s'", m["error"])
	}
}

func TestDeleteBook(t *testing.T) {
	clearTables()
	addBooks(1)

	req, _ := http.NewRequest("GET", "/books/1", nil)
	response := executeRequest(req)
	checkResponseCode(t, http.StatusOK, response.Code)

	req, _ = http.NewRequest("DELETE", "/books/1", nil)
	response = executeRequest(req)

	checkResponseCode(t, http.StatusOK, response.Code)

	req, _ = http.NewRequest("GET", "/books/1", nil)
	response = executeRequest(req)
	checkResponseCode(t, http.StatusNotFound, response.Code)
}

func TestDeleteBookNotThere(t *testing.T) {
	clearTables()

	req, _ := http.NewRequest("DELETE", "/books/1", nil)
	response := executeRequest(req)
	checkResponseCode(t, http.StatusOK, response.Code)
}

func addBooks(count int) {
	if count < 1 {
		count = 1
	}

	for i := 0; i < count; i++ {
		a.DB.Exec("INSERT INTO books(title, author, publisher, publish_date, rating, status) VALUES($1, $2, $3, $4, $5, $6)",
			"Book "+strconv.Itoa(i),
			"That One Guy",
			"Test Publisher",
			"2014-11-12T11:45:26.371Z",
			1,
			"CheckedIn")
	}
}
