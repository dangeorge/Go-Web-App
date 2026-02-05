package main

import (
	"html/template"
	"log"
	"net/http"
)

type PageData struct {
	Title string
}

func render(w http.ResponseWriter, page string, data PageData) {
	tmpl := template.Must(template.ParseFiles(
		"templates/base.html",
		"templates/" + page + ".html",
	))
	tmpl.ExecuteTemplate(w, "base", data)
}

func main() {
	mux := http.NewServeMux()

	mux.Handle("/static/", http.StripPrefix("/static/",
		http.FileServer(http.Dir("static")),
	))

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		render(w, "home", PageData{Title: "Home"})
	})

	mux.HandleFunc("/about", func(w http.ResponseWriter, r *http.Request) {
		render(w, "about", PageData{Title: "About"})
	})

	mux.HandleFunc("/contact", func(w http.ResponseWriter, r *http.Request) {
		render(w, "contact", PageData{Title: "Contact"})
	})

	log.Println("Server running on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", mux))
}
