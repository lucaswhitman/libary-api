package main_test

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"os"
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
