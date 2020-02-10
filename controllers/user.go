package controllers

import (
	"github.com/madjlzz/madlens/views"
	"net/http"
)

// NewUsers is used to create a new Users controller.
// This function will panic if templates are not
// parsed correctly, and should only be used during
// initial setup.
func NewUsers() *Users {
	return &Users{
		NewView: views.NewView("bootstrap", "views/users/new.gohtml"),
	}
}

//
func (u *Users) New(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	if err := u.NewView.Render(w, nil); err != nil {
		panic(err)
	}
}

type Users struct {
	NewView *views.View
}
