package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

type Book struct {
	ID     int    `json:id`
	Title  string `json:title`
	Author string `json:author`
	Year   string `json:year`
}

var books []Book

func main() {
	fmt.Println("hello")

	books := createData()

	fmt.Println(books)

	router := mux.NewRouter()
	router.HandleFunc("/books", getBooks).Methods("GET")
	router.HandleFunc("/book/{id}", getBook).Methods("GET")
	router.HandleFunc("/books", addBook).Methods("POST")
	router.HandleFunc("/books", updateBook).Methods("PUT")
	router.HandleFunc("/book/{id}", removeBook).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8080", router))
}

func createData() []Book {
	bookOne := Book{ID: 1, Title: "Golang Pointers", Author: "Mr. Golang", Year: "2010"}
	bookTwo := Book{ID: 2, Title: "Goroutines", Author: "Mr. Goroutines", Year: "2011"}
	bookThree := Book{ID: 3, Title: "Golang Routers", Author: "Mr. Router", Year: "2012"}
	bookFour := Book{ID: 4, Title: "Golang Concurrency", Author: "Mr. Concurrency", Year: "2013"}
	bookFive := Book{ID: 5, Title: "Golang Good Parts", Author: "Mr. Good", Year: "2014"}
	books := append(books, bookOne, bookTwo, bookThree, bookFour, bookFive)
	return books
}

func getBooks(w http.ResponseWriter, r *http.Request) {
	log.Println("getBooks")
}

func getBook(w http.ResponseWriter, r *http.Request) {
	log.Println("getBook")
}

func addBook(w http.ResponseWriter, r *http.Request) {
	log.Println("addBook")
}

func updateBook(w http.ResponseWriter, r *http.Request) {
	log.Println("updateBook")
}

func removeBook(w http.ResponseWriter, r *http.Request) {
	log.Println("removeBook")
}
