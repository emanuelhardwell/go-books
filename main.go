package main

import (
	"database/sql"
	"fmt"
	"go-books/internal/service"
	"go-books/internal/store"
	"go-books/internal/transport"
	"log"
	"net/http"

	_ "github.com/mattn/go-sqlite3"
)

func main() {

	// Create conecction to DB
	db, err := sql.Open("sqlite3", "./books.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Create table if not exist
	query := "create table if not exists Books (id integer primary key autoincrement, title text not null, author text not null)"
	if _, err := db.Exec(query); err != nil {
		log.Fatal(err)
	}

	// inject our dependencies
	bookStore :=
		store.New(db)

	bookService :=
		service.New(bookStore)

	bookHandler := transport.New(bookService)

	// config the routes
	http.HandleFunc("/books", bookHandler.HandleBooksGetPost)
	http.HandleFunc("/books/", bookHandler.HandlerBookGetDeleteById)

	fmt.Println("º Servidor ejecutándose en http://localhost:8080")
	fmt.Println("= API Endpoints:")

	// start and listen the server
	log.Fatal(http.ListenAndServe(":8080", nil))
}
