package main

import (
	"html/template"
	"net/http"

	"github.com/gorilla/mux"
)

var templates *template.Template

func main() {
	templates = template.Must(template.ParseGlob("templates/*.html"))
	r := mux.NewRouter()
	r.HandleFunc("/", indexHandler).Methods("GET")

	http.Handle("/", r)
	http.ListenAndServe(":5000", nil)
}

func indexHandler(w http.ResponseWriter, r *http.Request) {

	templates.ExecuteTemplate(w, "index.html", nil)
}
