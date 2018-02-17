package main_test

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

	"."
)

var a main.App

type configuration struct {
	Host         string `json:"host"`
	Port         int    `json:"port"`
	Username     string `json:"username"`
	Password     string `json:"password"`
	DatabaseName string `json:"databaseName"`
}

func TestMain(m *testing.M) {
	a = main.App{}

	file, _ := os.Open("config_test.json")
	decoder := json.NewDecoder(file)
	conf := configuration{}
	err := decoder.Decode(&conf)
	if err != nil {
		log.Fatal(err)
	}

	a.Initialize(conf.Host, conf.Port, conf.Username, conf.Password, conf.DatabaseName)

	ensureTablesExists()

	code := m.Run()

	clearTables()

	os.Exit(code)
}

func ensureTablesExists() {
	booksTableCreationQuery, err := ioutil.ReadFile("sql/createBooksTable.sql") // just pass the file name
	if err != nil {
		log.Fatal(err)
	}
	if _, err := a.DB.Exec(string(booksTableCreationQuery)); err != nil {
		log.Fatal(err)
	}

	authorsTableCreationQuery, err := ioutil.ReadFile("sql/createAuthorsTable.sql") // just pass the file name
	if err != nil {
		log.Fatal(err)
	}
	if _, err := a.DB.Exec(string(authorsTableCreationQuery)); err != nil {
		log.Fatal(err)
	}
}

func clearTables() {
	a.DB.Exec("DELETE FROM books")
	a.DB.Exec("DELETE FROM authors")
	a.DB.Exec("ALTER SEQUENCE books_id_seq RESTART WITH 1")
	a.DB.Exec("ALTER SEQUENCE authors_id_seq RESTART WITH 1")
}

func TestEmptyTable(t *testing.T) {
	clearTables()

	req, _ := http.NewRequest("GET", "/products", nil)
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

func TestGetNonExistentBook(t *testing.T) {
	clearTables()

	req, _ := http.NewRequest("GET", "/books/11", nil)
	response := executeRequest(req)

	checkResponseCode(t, http.StatusNotFound, response.Code)

	var m map[string]string
	json.Unmarshal(response.Body.Bytes(), &m)
	if m["error"] != "Product not found" {
		t.Errorf("Expected the 'error' key of the response to be set to 'Book not found'. Got '%s'", m["error"])
	}
}

func TestCreateBook(t *testing.T) {
	clearTables()

	payload := []byte(`{
		"title":"Hitchhikers Guide to the Galaxy",
		"author":{
			"firstName": "Douglas", 
			"lastName":"Adams"}
		},
		"publisher": "Pan Books",
		"publishDate": "10-12-1979",
		"rating": 3,
		"status": "CheckedOut"`)

	req, _ := http.NewRequest("POST", "/book", bytes.NewBuffer(payload))
	response := executeRequest(req)

	checkResponseCode(t, http.StatusCreated, response.Code)

	var m map[string]interface{}
	json.Unmarshal(response.Body.Bytes(), &m)

	if m["title"] != "Hitchhikers Guide to the Galaxy" {
		t.Errorf("Expected product name to be 'Hitchhikers Guide to the Galaxy'. Got '%v'", m["name"])
	}

	//todo: more checks

	// the id is compared to 1.0 because JSON unmarshaling converts numbers to
	// floats, when the target is a map[string]interface{}
	/*if m["id"] != 1.0 {
		t.Errorf("Expected product ID to be '1'. Got '%v'", m["id"])
	}
	*/
}

func TestGetBook(t *testing.T) {
	clearTables()
	addBooks(1)

	req, _ := http.NewRequest("GET", "/book/1", nil)
	response := executeRequest(req)

	checkResponseCode(t, http.StatusOK, response.Code)
}

func addBooks(count int) {
	if count < 1 {
		count = 1
	}

	//insert one author
	a.DB.Exec("INSERT INTO authors(firstname, lastname) VALUES($1, $2)")

	for i := 0; i < count; i++ {
		a.DB.Exec("INSERT INTO books(title, author_id, publisher, publish_date, rating, status) VALUES($1, $2, $3, $4, $5, $6)",
			"Book "+strconv.Itoa(i),
			1,
			"Test Publisher",
			"2011-01-01",
			3,
			"CheckedIn")
	}
}

func TestUpdateBook(t *testing.T) {
	clearTables()
	addBooks(1)

	req, _ := http.NewRequest("GET", "/book/1", nil)
	response := executeRequest(req)
	var originalBook map[string]interface{}
	json.Unmarshal(response.Body.Bytes(), &originalBook)

	payload := []byte(`{
		"title":"Hitchhikers Guide to the Galaxy",
		"author":{
			"firstName": "Douglas", 
			"lastName":"Adams"}
		},
		"publisher": "Pan Books",
		"publishDate": "10-12-1979",
		"rating": 3,
		"status": "CheckedOut"`)

	req, _ = http.NewRequest("PUT", "/book/1", bytes.NewBuffer(payload))
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

	if m["publisher"] == originalBook["publisher"] {
		t.Errorf("Expected the publisher to change from '%v' to '%v'. Got '%v'", originalBook["publisher"], m["publisher"], m["publisher"])
	}
}

func TestDeleteBook(t *testing.T) {
	clearTables()
	addBooks(1)

	req, _ := http.NewRequest("GET", "/book/1", nil)
	response := executeRequest(req)
	checkResponseCode(t, http.StatusOK, response.Code)

	req, _ = http.NewRequest("DELETE", "/book/1", nil)
	response = executeRequest(req)

	checkResponseCode(t, http.StatusOK, response.Code)

	req, _ = http.NewRequest("GET", "/book/1", nil)
	response = executeRequest(req)
	checkResponseCode(t, http.StatusNotFound, response.Code)
}
