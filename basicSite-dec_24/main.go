package main

import (
	"fmt"
	"net/http"
	"text/template"
)

// template initialization
func templateInit() *template.Template {
	templates := template.Must(template.ParseGlob("templates/*.html"))

	return templates
}

// TODO: handler functions
func buildHandler(fn func(http.ResponseWriter, *http.Request, *template.Template)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		t := templateInit()
		fn(w, r, t)
	}
}

// TODO: I want to be able to render a template based on the request parsed via regex
func statucPageHandler(w http.ResponseWriter, r *http.Request, t *template.Template) {
	fmt.Printf("r.URL.Path: %v\n", r.URL.Path)
	t.ExecuteTemplate(w, r.URL.Path[1:]+".html", nil)
}

// server initialization, creates the mux and adds our handlers
func serverInit() *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("/", buildHandler(statucPageHandler))

	return mux
}

func main() {
	mux := serverInit()
	http.ListenAndServe("localhost:8008", mux)
}
