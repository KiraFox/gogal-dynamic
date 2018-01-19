package controllers

import (
	"fmt"
	"net/http"

	"github.com/KiraFox/gogal-dynamic/views"
)

// This is the controller for the "users" resource
type Users struct {
	//Update controller to store a "new" user view template
	NewView *views.View
}

// This function handles the logic for parsing "new" users view template and
// returning the information to the Users controller struct
func NewUsers() *Users {
	return &Users{
		NewView: views.NewView("bootstrap", "views/users/new.gohtml"),
	}
}

// This method is used to render the view stored in the NewView field of the
// Users controller struct for the Sign Up ("new" users) webpage
// GET /signup
func (u *Users) New(w http.ResponseWriter, r *http.Request) {
	if err := u.NewView.Render(w, nil); err != nil {
		panic(err)
	}
}

// This method is used to process the signup form when a user tries to create
// a new user account.
// POST /signup
func (u *Users) Create(w http.ResponseWriter, r *http.Request) {
	// When form is submitted to your web application it will be included in
	// part of the http.Request parameter (r) that is passed into your handler
	// function. Therefor we can call a method ParseForm on the request to parse
	// the information that is then stored in a map. We retrieve the information
	// in the map using the ["key"] syntax and the keys are based on the "name"
	// we have set in the new users template in the form HTML code.
	if err := r.ParseForm(); err != nil {
		panic(err)
	}
	fmt.Fprintln(w, r.PostForm["email"])
	fmt.Fprintln(w, r.PostForm["password"])
}
