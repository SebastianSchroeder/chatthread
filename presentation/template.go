package presentation

import (
	"html/template"
	"net/http"
)

var templates = template.Must(template.ParseFiles(
	"templates/page.html",
))

func renderPage(w http.ResponseWriter, template string, data interface{}) {
	err := templates.ExecuteTemplate(w, template, data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
