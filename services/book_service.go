package services

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/lucaswhitman/library-api/models"
)

type BookService struct {
	DB *sql.DB
}

func (bookService *BookService) GetBook(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		RespondWithError(w, http.StatusBadRequest, "Invalid book ID")
		return
	}
	b := models.Book{ID: id}
	if err := b.GetBook(bookService.DB); err != nil {
		switch err {
		case sql.ErrNoRows:
			RespondWithError(w, http.StatusNotFound, "Book not found")
		default:
			RespondWithError(w, http.StatusInternalServerError, err.Error())
		}
		return
	}

	RespondWithJSON(w, http.StatusOK, b)
}

func (bookService *BookService) GetBooks(w http.ResponseWriter, r *http.Request) {
	count, _ := strconv.Atoi(r.FormValue("count"))
	start, _ := strconv.Atoi(r.FormValue("start"))

	if count > 10 || count < 1 {
		count = 10
	}
	if start < 0 {
		start = 0
	}

	books, err := models.GetBooks(bookService.DB, start, count)
	if err != nil {
		RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	RespondWithJSON(w, http.StatusOK, books)
}

func (bookService *BookService) CreateBook(w http.ResponseWriter, r *http.Request) {
	var b models.Book
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&b); err != nil {
		RespondWithError(w, http.StatusBadRequest, err.Error())
		return
	}
	defer r.Body.Close()

	if err := b.CreateBook(bookService.DB); err != nil {
		RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	RespondWithJSON(w, http.StatusCreated, b)
}

func (bookService *BookService) UpdateBook(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		RespondWithError(w, http.StatusBadRequest, "Invalid book ID")
		return
	}

	var b models.Book
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&b); err != nil {
		RespondWithError(w, http.StatusBadRequest, err.Error())
		return
	}
	defer r.Body.Close()
	b.ID = id

	if err := b.UpdateBook(bookService.DB); err != nil {
		RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	RespondWithJSON(w, http.StatusOK, b)
}

func (bookService *BookService) DeleteBook(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		RespondWithError(w, http.StatusBadRequest, "Invalid Book ID")
		return
	}

	b := models.Book{ID: id}
	if err := b.DeleteBook(bookService.DB); err != nil {
		RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	RespondWithJSON(w, http.StatusOK, map[string]string{"result": "success"})
}

func RespondWithError(w http.ResponseWriter, code int, message string) {
	RespondWithJSON(w, code, map[string]string{"error": message})
}

func RespondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}
