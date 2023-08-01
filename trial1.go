package main

import (
	"fmt"
	"log"
	"net/http"
)

func Hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello World")
}

func main() {
	http.HandleFunc("/", Hello)

	log.Fatal(http.ListenAndServe(":5100", nil))
}
