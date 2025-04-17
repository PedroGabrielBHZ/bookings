package render

import (
	"app/pkg/models"
	"html/template"
	"log"
	"net/http"
	"path/filepath"
)

var templateCache = make(map[string]*template.Template)

func RenderTemplate(w http.ResponseWriter, templateName string, data *models.TemplateData) {
	tmpl, err := getTemplate(templateName)
	if err != nil {
		log.Printf("Error getting template: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	if err := tmpl.Execute(w, data); err != nil {
		log.Printf("Error executing template: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}

func getTemplate(templateName string) (*template.Template, error) {
	if tmpl, exists := templateCache[templateName]; exists {
		log.Println("Using cached template")
		return tmpl, nil
	}
	tmpl, err := CreateTemplate(templateName)
	if err != nil {
		return nil, err
	}
	templateCache[templateName] = tmpl
	return tmpl, nil
}

func CreateTemplate(templateName string) (*template.Template, error) {
	baseLayouts, err := filepath.Glob(filepath.Join("templates", "*.layout.tmpl.html"))
	if err != nil {
		log.Printf("Error finding base layouts: %v", err)
		return nil, err
	}

	templates := []string{filepath.Join("templates", templateName+".html")}
	templates = append(templates, baseLayouts...)

	tmpl, err := template.ParseFiles(templates...)
	if err != nil {
		log.Printf("Error creating template: %v", err)
		return nil, err
	}
	return tmpl, nil
}
