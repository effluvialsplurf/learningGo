package main

import (
	"net/http"
	"text/template"

	"basicSite/pages"
)

// TODO: handler functions
func buildHandler(fn func(http.ResponseWriter, *http.Request, *template.Template)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		t := pages.TemplateInit("pages/templates")
		fn(w, r, t[0])
	}
}

// TODO: I want to be able to render a template based on the request parsed via regex
func staticPageHandler(w http.ResponseWriter, r *http.Request, t *template.Template) {
	t.ExecuteTemplate(w, r.URL.Path[1:]+".html", nil) // we have to remove the leading slash from the path
}

// server initialization, creates the mux and adds our handlers
func serverInit() *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("/", buildHandler(staticPageHandler))

	return mux
}

func main() {
	mux := serverInit()
	http.ListenAndServe("localhost:8008", mux)
}
