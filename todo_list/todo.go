package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

var tmpl *template.Template

type Todo struct {
	Item string
	Done bool
}

type PageData struct {
	Title string
	Todos []Todo
}

func todo(w http.ResponseWriter, r *http.Request) {

	data := PageData{
		Title: "Todo List",
		Todos: []Todo{
			{
				Item: "Read", Done: true,
			},
			{
				Item: "Write", Done: false,
			},
			{
				Item: "Breakfast", Done: false,
			},
		},
	}

	tmpl.Execute(w, data)
}

func main() {
	fmt.Println("Hello")

	mux := http.NewServeMux()
	tmpl = template.Must(template.ParseFiles("./templates/index.html"))

	fs := http.FileServer(http.Dir("./static"))
	mux.Handle("/static/", http.StripPrefix("/static/", fs))
	mux.HandleFunc("/todo", todo)

	log.Fatal(http.ListenAndServe(":9000", mux))
}
