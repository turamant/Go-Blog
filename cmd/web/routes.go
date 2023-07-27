package main

import (
	"net/http"
	"github.com/bmizerany/pat"

	"github.com/justinas/alice"
	
)

func (app *application) routes() http.Handler {

	standartMiddleware := alice.New(app.recoverPanic, app.logRequest, secureHeaders)
	dynamicMiddleware := alice.New(app.session.Enable)

	mux := pat.New()
	mux.Get("/", dynamicMiddleware.ThenFunc(app.home))
	mux.Get("/post/create", dynamicMiddleware.ThenFunc(app.createPostForm))
	mux.Post("/post/create", dynamicMiddleware.ThenFunc(app.createPost))
	mux.Get("/post/:id", dynamicMiddleware.ThenFunc(app.showPost))		
	
	fileServer := http.FileServer(http.Dir("./ui/static"))
	mux.Get("/static/", http.StripPrefix("/static", fileServer))
	
	return standartMiddleware.Then(mux)
}
