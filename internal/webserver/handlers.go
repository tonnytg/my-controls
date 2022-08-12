package webserver

import (
	"html/template"
	"log"
	"net/http"
)

var (
	tmpl = template.Must(template.ParseGlob("internal/webserver/templates/*.tmpl"))
	fs   = http.FileServer(http.Dir("./internal/webserver/templates/static"))
)

type PerimeterType struct {
	ID           int
	Title        string
	Type         string
	ProjectCount int
}

func IndexHandler(w http.ResponseWriter, r *http.Request) {

	err := tmpl.ExecuteTemplate(w, "Index", []PerimeterType{
		PerimeterType{1, "Perimeter One", "Regular", 10},
		PerimeterType{2, "Perimeter Two", "Regular", 20},
		PerimeterType{3, "Perimeter Three", "Regular", 30},
	})
	if err != nil {
		log.Printf("Error executing template: %v", err)
		return
	}
}

func LoadHandlers() {

	http.HandleFunc("/", IndexHandler)
	http.Handle("/static/", http.StripPrefix("/static", fs))
}
