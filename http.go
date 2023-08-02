package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Book struct {
	Title  string `json: "title"`
	Author string `json: "author"`
	Pages  int    `json: "pages"`
}

func GetBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Conten-Type", "application/json")

	book := Book{
		Title:  "Shiva",
		Author: "Amish",
		Pages:  450,
	}

	json.NewEncoder(w).Encode(book)
}

func Hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello World")
}

func main() {
	http.HandleFunc("/", Hello)
	http.HandleFunc("/book", GetBook)

	log.Fatal(http.ListenAndServe(":5100", nil))
}
