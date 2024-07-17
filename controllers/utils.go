package controllers

import (
	"net/http"
	"path/filepath"
	"text/template"
)

const codeRedirect = 301

var temp = template.Must(template.ParseGlob("templates/**/*.html"))

func renderTemplate(w http.ResponseWriter, tmpl string, data any) {
	// Obtém o nome base do template
	baseTemplate := "base.html"
	baseTemplatePath := filepath.Join("templates", baseTemplate)
	tmplPath := filepath.Join("templates", tmpl)
	partialsPath := filepath.Join("templates/layout/*.html")

	t, err := template.ParseFiles(baseTemplatePath, tmplPath)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	t, err = t.ParseGlob(partialsPath)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = t.ExecuteTemplate(w, baseTemplate, data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	// Combina o template base com o template específico
	//err := template.Must(template.ParseFiles(filepath.Join("templates", baseTemplate), filepath.Join("templates", tmpl))).ExecuteTemplate(w, baseTemplate, data)
}
