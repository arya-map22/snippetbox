package main

import (
	"net/http"
)

func (app *application) routes() *http.ServeMux {
	mux := http.NewServeMux()
	fileServer := http.FileServer(neuteredFileSystem{fs: http.Dir(`D:\web_dev\go\snippetbox\ui\static\`)})

	mux.Handle("GET /static/", http.StripPrefix("/static", fileServer))

	mux.HandleFunc("GET /{$}", app.home)
	mux.HandleFunc("GET /snippet/view/{id}", app.snippetView)
	mux.HandleFunc("GET /snippet/create", app.snippetCreate)

	mux.HandleFunc("POST /snippet/create", app.snippetCreatePost)

	return mux
}
