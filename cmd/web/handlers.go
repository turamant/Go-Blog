package main

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/turamant/Go-Blog/pkg/forms"
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
	
	app.render(w, r, "show.page.html", &templateData{
		Post: p,
	})

}

func (app *application) createPost(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		app.clientError(w, http.StatusBadRequest)
		return
	}
	form := forms.New(r.PostForm)
	form.Required("title", "content", "expires")
	form.MaxLength("title", 100)
	form.PermittedValues("expires", "365", "7","1")

	if !form.Valid() {
		app.render(w, r, "create.page.html", &templateData{Form: form})
		return
	}

	id, err := app.posts.Insert(form.Get("title"), form.Get("content"), form.Get("expires"))
	if err != nil {
		app.serverError(w, err)
		return
	}
	app.session.Put(r,"flash", "Post successfully created!")
	http.Redirect(w, r, fmt.Sprintf("/post/%d", id), http.StatusSeeOther)
}

func (app *application) createPostForm(w http.ResponseWriter, r *http.Request) {
	app.render(w, r, "create.page.html", &templateData{
		Form: forms.New(nil),
	})
}


func (app *application) signupUserForm(w http.ResponseWriter, r *http.Request) {
	app.render(w, r, "signup.page.html", &templateData{
		Form: forms.New(nil),
		})
}

func (app *application) signupUser(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		app.clientError(w, http.StatusBadRequest)
		return
	}
	
	form := forms.New(r.PostForm)
	form.Required("name", "email", "password")
	form.MaxLength("name", 255)
	form.MaxLength("email", 255)
	form.MatchesPattern("email", forms.EmailRX)
	form.MinLength("password", 10)
	
	if !form.Valid() {
		app.render(w, r, "signup.page.html", &templateData{Form: form})
		return
	}
	err = app.users.Insert(form.Get("name"), form.Get("email"), form.Get("password"))
	if err != nil{
		if errors.Is(err, models.ErrDuplicateEmail) {
			form.Errors.Add("email", "Address is already in use")
			app.render(w, r, "signup.page.html", &templateData{Form: form})
		} else {
			app.serverError(w, err)
		}
		return
	}
	app.session.Put(r, "flash", "Your signup was successful. Please log in.")
	http.Redirect(w, r, "/user/login", http.StatusSeeOther)
}	


func (app *application) loginUserForm(w http.ResponseWriter, r *http.Request) {
	app.render(w, r, "login.page.html", &templateData{
		Form: forms.New(nil),
		})
}

func (app *application) loginUser(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Authenticate and login the user...")
}

func (app *application) logoutUser(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Logout the user...")
}

