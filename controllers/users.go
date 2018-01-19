package controllers

import (
	"fmt"
	"net/http"

	"github.com/KiraFox/gogal-dynamic/views"

	//Using this package to help parse the signup form
	"github.com/gorilla/schema"
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

// This method is used to process the Signup form when a user tries to create
// a new user account.
// POST /signup
func (u *Users) Create(w http.ResponseWriter, r *http.Request) {
	// Need to parse the form that is POSTed as part of our http.Request when the
	// Signup form is submitted; otherwise the form is ignored when submitted
	// and we have no values to work with.
	if err := r.ParseForm(); err != nil {
		panic(err)
	}
	// Intialize a decoder from the schema package to use on the parsed form.
	dec := schema.NewDecoder()
	// Intialize a SignupForm so we have a destination for the parsed and decoded
	// form information.
	form := SignupForm{}
	// Run the decoder using the Decode method and give it the destination (&form)
	// where you want the decoded parsed form information stored and the source
	// (r.PostForm) you want to decode to begin with.
	if err := dec.Decode(&form, r.PostForm); err != nil {
		panic(err)
	}

	fmt.Fprintln(w, form)
}

// This is the struct to hold the information submitted using the Signup form
type SignupForm struct {
	// Use struct tags so gorilla/schema package knows about the input fields
	// in the Signup form.
	Email    string `schema:"email"`
	Password string `schema:"password"`
}
