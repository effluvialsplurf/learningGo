package pages

import "text/template"

type Page struct {
	Title string
	Body  []byte
}

// template initialization
func TemplateInit(dir string) []*template.Template {
	staticTemplates := template.Must(template.ParseGlob(dir + "/*.html"))

	return []*template.Template{staticTemplates}
}
