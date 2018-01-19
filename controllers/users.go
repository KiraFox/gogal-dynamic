package controllers

import (
	"fmt"
	"net/http"

	"github.com/KiraFox/gogal-dynamic/views"
	//Using this package to help parse the signup form
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
	// Intialize the destination variable
	var form SignupForm

	// Call parseForm helper function and pass it the http.Request to parse &
	// decode and the destination for the final information, and check for any
	// returned errors
	if err := parseForm(r, &form); err != nil {
		panic(err)
	}
	// Print to check we can access the fields in SignupForm and that it has the
	// correct values associated.
	fmt.Fprintln(w, "Email is", form.Email)
	fmt.Fprintln(w, "Password is", form.Password)
}

// This is the struct to hold the information submitted using the Signup form
type SignupForm struct {
	// Use struct tags so gorilla/schema package knows about the input fields
	// in the Signup form.
	Email    string `schema:"email"`
	Password string `schema:"password"`
}
