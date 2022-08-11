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

func IndexHandler(w http.ResponseWriter, r *http.Request) {

	err := tmpl.ExecuteTemplate(w, "Index", nil)
	if err != nil {
		log.Printf("Error executing template: %v", err)
		return
	}
}

func LoadHandlers() {

	http.HandleFunc("/", IndexHandler)
	http.Handle("/static/", http.StripPrefix("/static", fs))
}
