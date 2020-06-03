package main

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/oniric85/snippetbox/pkg/models"
)

func (app *application) home(w http.ResponseWriter, r *http.Request) {
	s, err := app.snippets.Latest()

	if err != nil {
		app.serverError(w, err)
		return
	}

	data := &templateData{
		Snippets: s,
	}

	app.render(w, r, "home.page.tmpl", data)
}

func (app *application) showSnippet(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get(":id"))
	if err != nil || id < 1 {
		app.notFound(w)
		return
	}

	s, err := app.snippets.Get(id)

	if err == models.ErrNoRecord {
		app.notFound(w)
		return
	}

	if err != nil {
		app.serverError(w, err)
		return
	}

	data := &templateData{
		Snippet: s,
	}

	app.render(w, r, "show.page.tmpl", data)
}

func (app *application) createSnippetForm(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Create a new snippet..."))
}

func (app *application) createSnippet(w http.ResponseWriter, r *http.Request) {
	title := "O snail"
	content := "O snail\nClimb Mount Fuji,\nBut slowly, slowly!"
	expires := "7"

	id, err := app.snippets.Insert(title, content, expires)

	if err != nil {
		app.serverError(w, err)
		return
	}

	http.Redirect(w, r, fmt.Sprintf("/snippet?id=%d", id), http.StatusSeeOther)
}
