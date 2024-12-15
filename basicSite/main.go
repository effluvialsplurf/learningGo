package main

import (
	"io/fs"
	"net/http"
	"text/template"

	"basicSite/pages"
	"basicSite/viteapp"
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

func appHandler(w http.ResponseWriter, r *http.Request) {
	apDir, err := fs.Sub(viteapp.App, "canvasGames/dist")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	http.StripPrefix("/app/", http.FileServerFS(apDir)).ServeHTTP(w, r)
}

// server initialization, creates the mux and adds our handlers
func serverInit() *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("/app/", appHandler)
	mux.HandleFunc("/", buildHandler(staticPageHandler))

	return mux
}

func main() {
	mux := serverInit()
	http.ListenAndServe("localhost:8008", mux)
}
