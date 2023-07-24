package main

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"

)

func (app *application) home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		app.notFound(w)
		return
	}

	files := []string {
		"./ui/html/home.page.html",
		"./ui/html/base.layout.html",
		"./ui/html/footer.partial.html",
	}

	ts, err := template.ParseFiles(files...)
	if err != nil {
		app.serverError(w, err)
		return
	}
	err = ts.Execute(w, nil)
	if err != nil {
		app.serverError(w, err)
	}
}

func (app *application) showPost(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil || id < 1{
		app.notFound(w)
		return

	}
	fmt.Fprintf(w, "This is specific post id: %d", id)
}

func (app *application) createPost(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost{
		w.Header().Set("Allow", http.MethodPost)
		//w.WriteHeader(405)
		//w.Write([]byte("Method Not Allowed"))
		app.clientError(w, http.StatusMethodNotAllowed)
		return
	}
	w.Write([]byte("This is create post page"))
}