package app

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
	"github.com/lucaswhitman/library-api/services"
)

type App struct {
	Router      *mux.Router
	DB          *sql.DB
	bookService services.BookService
}

func (a *App) Initialize(host string, port int, user string, password string, dbname string) {
	// todo: in production, we'd enable SSL
	connectionString := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	var err error
	a.DB, err = sql.Open("postgres", connectionString)
	if err != nil {
		log.Fatal(err)
	}

	a.Router = mux.NewRouter()
	a.initializeRoutes()

	a.bookService = services.BookService{DB: a.DB}

}

func (a *App) Run(port int) {
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), a.Router))
}

func (a *App) initializeRoutes() {
	a.Router.HandleFunc("/books", a.bookService.GetBooks).Methods("GET")
	a.Router.HandleFunc("/books", a.bookService.CreateBook).Methods("POST")
	a.Router.HandleFunc("/books/{id:[0-9]+}", a.bookService.GetBook).Methods("GET")
	a.Router.HandleFunc("/books/{id:[0-9]+}", a.bookService.UpdateBook).Methods("PUT")
	a.Router.HandleFunc("/books/{id:[0-9]+}", a.bookService.DeleteBook).Methods("DELETE")
}
