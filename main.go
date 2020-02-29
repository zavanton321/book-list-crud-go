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

	router := mux.NewRouter()
	router.HandleFunc("/books", getBooks).Methods("GET")
	router.HandleFunc("book/{id}", getBook).Methods("GET")
	router.HandleFunc("/books", addBook).Methods("POST")
	router.HandleFunc("/books", updateBook).Methods("PUT")
	router.HandleFunc("/book/{id}", removeBook).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8080", router))
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
