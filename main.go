package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"strconv"
)

type Book struct {
	ID     int    `json:id`
	Title  string `json:title`
	Author string `json:author`
	Year   string `json:year`
}

var books []Book

func main() {
	books = createData()

	router := mux.NewRouter()
	router.HandleFunc("/books", getBooks).Methods("GET")
	router.HandleFunc("/books/{id}", getBook).Methods("GET")
	router.HandleFunc("/books", addBook).Methods("POST")
	router.HandleFunc("/books", updateBook).Methods("PUT")
	router.HandleFunc("/books/{id}", removeBook).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8080", router))
}

func getBooks(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(books)
}

func getBook(w http.ResponseWriter, r *http.Request) {
	log.Println("getBook")

	params := mux.Vars(r)
	idString := params["id"]
	targetId, err := strconv.Atoi(idString)

	if err != nil {
		log.Println("Failed to parse idString!", err)
		return
	}

	for _, book := range books {
		if book.ID == targetId {
			log.Println("The requested book: ", book)
			json.NewEncoder(w).Encode(&book)
		}
	}
}

func addBook(w http.ResponseWriter, r *http.Request) {
	log.Println("addBook")

	var book Book
	json.NewDecoder(r.Body).Decode(&book)
	books = append(books, book)
	log.Println("New book: ", book)

	json.NewEncoder(w).Encode(books)
}

func updateBook(w http.ResponseWriter, r *http.Request) {
	log.Println("updateBook")
}

func removeBook(w http.ResponseWriter, r *http.Request) {
	log.Println("removeBook")
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
