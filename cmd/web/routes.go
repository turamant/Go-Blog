package main

import (
	"net/http"

	"github.com/justinas/alice"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/v5/middleware"
)

func (app *application) routes() http.Handler {

	standartMiddleware := alice.New(app.recoverPanic, app.logRequest, secureHeaders)

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Get("/", app.home)
	r.Get("/post", app.showPost)
	r.Post("/post/create", app.createPost)

	fileServer := http.FileServer(http.Dir("./ui/static"))
	r.Handle("/static/", http.StripPrefix("/static", fileServer))
	
	return standartMiddleware.Then(r)
}
