package main

import (
	"net/http"
	"text/template"
)

// TODO: handler functions
func buildHandler(fn func(http.ResponseWriter, *http.Request, *template.Template)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		t := pages.templateInit()
		fn(w, r, t)
	}
}

// TODO: I want to be able to render a template based on the request parsed via regex
func statucPageHandler(w http.ResponseWriter, r *http.Request, t *template.Template) {
	t.ExecuteTemplate(w, r.URL.Path[1:]+".html", nil) // we have to remove the leading slash from the path
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
