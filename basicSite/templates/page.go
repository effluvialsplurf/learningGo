package pages

import "text/template"

type Page struct {
	Title string
	Body  []byte
}

// template initialization
func templateInit() []*template.Template {
	staticTemplates := template.Must(template.ParseGlob("templates/*.html"))

	return []*template.Template{staticTemplates}
}
