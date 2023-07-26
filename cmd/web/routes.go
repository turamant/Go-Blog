package main

import (
	"net/http"
	"github.com/bmizerany/pat"

	"github.com/justinas/alice"
	
)

func (app *application) routes() http.Handler {

	standartMiddleware := alice.New(app.recoverPanic, app.logRequest, secureHeaders)

	mux := pat.New()
	mux.Get("/", http.HandlerFunc(app.home))
	mux.Get("/post/create", http.HandlerFunc(app.createPostForm))
	mux.Post("/post/create", http.HandlerFunc(app.createPost))
	mux.Get("/post/:id", http.HandlerFunc(app.showPost))		
	
	

	fileServer := http.FileServer(http.Dir("./ui/static"))
	mux.Get("/static/", http.StripPrefix("/static", fileServer))
	
	return standartMiddleware.Then(mux)
}
