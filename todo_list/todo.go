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
				Item: "Good Morning", Done: false,
			},
			{
				Item: "Brush", Done: false,
			},
			{
				Item: "Bath", Done: false,
			},
			{
				Item: "Breakfast", Done: false,
			},
			{
				Item: "2 hour study(learn different technology) session", Done: false,
			},
			{
				Item: "30 min break", Done: false,
			},
			{
				Item: "1.5 hour coding session", Done: false,
			},
			{
				Item: "1.5 hour Lunch Break and nap time", Done: false,
			},
			{
				Item: "30 min reading and coffee time", Done: false,
			},
			{
				Item: "2 hour blog case study and write some part of it and make git commits", Done: false,
			},
			{
				Item: "2 hour dinner break and netflix", Done: false,
			},
			{
				Item: "1.5 hour college assignments and code completion", Done: false,
			},
			{
				Item: "Good Night", Done: false,
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
