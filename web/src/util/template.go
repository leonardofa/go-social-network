package util

import (
	"html/template"
	"net/http"
)

var templates *template.Template

// LoadHtmlTemplates loads the HTML templates from the "view" directory.
func LoadHtmlTemplates() {
	templates = template.Must(template.ParseGlob("view/*.html"))
}

// RenderTemplate renders the specified HTML template with the given data.
func RenderTemplate(w http.ResponseWriter, template string, data any) {
	err := templates.ExecuteTemplate(w, template, data)
	if err != nil {
		return
	}
}
