package main

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/turamant/Go-Blog/pkg/models"

)


func (app *application) home(w http.ResponseWriter, r *http.Request) {
	// if r.URL.Path != "/" {
	// 	app.notFound(w)
	// 	return
	// }

	p, err := app.posts.Latest()
	if err != nil {
		app.serverError(w, err)
		return
	}
	app.render(w, r, "home.page.html", &templateData{Posts: p})
}

func (app *application) showPost(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get(":id"))
	
	if err != nil || id < 1 {
		fmt.Printf("-?????--%d----%s",id, err)
		app.notFound(w)
		return
	}
	p, err := app.posts.Get(id)
	if err != nil {
		if errors.Is(err, models.ErrNoRecord) {
			app.notFound(w)
		} else {
			app.serverError(w, err)
		}
		return
	}
	app.render(w, r, "show.page.html", &templateData{Post: p})

}

func (app *application) createPost(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		app.clientError(w, http.StatusBadRequest)
		return
	}

	title := r.PostForm.Get("title")
	content := r.PostForm.Get("content")
	expires := r.PostForm.Get("expires")

	id, err := app.posts.Insert(title, content, expires)
	if err != nil {
		app.serverError(w, err)
		return
	}
	http.Redirect(w, r, fmt.Sprintf("/post/%d", id), http.StatusSeeOther)
}

func (app *application) createPostForm(w http.ResponseWriter, r *http.Request) {
	app.render(w, r, "create.page.html", nil)
}

