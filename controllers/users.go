package controllers

import (
	"fmt"
	"net/http"

	"github.com/KiraFox/gogal-dynamic/models"
	"github.com/KiraFox/gogal-dynamic/views"
)

// This function handles the logic for parsing "new" users view template and
// returning the information to the Users controller struct. It also takes a
// user service and assigns it to the same struct to be used by its methods.
func NewUsers(us *models.UserService) *Users {
	return &Users{
		NewView: views.NewView("bootstrap", "users/new"),
		us:      us,
	}
}

// This is the controller for the "users" resource
type Users struct {
	// Store a "new" user view template
	NewView *views.View
	// Store a UserService instance for handlers to access
	us *models.UserService
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

	// Use the information from the Signup Form to fill out information in the
	// User model
	user := models.User{
		Email:    form.Email,
		Name:     form.Name,
		Password: form.Password,
	}

	// Use the now filled out User model struct and run its Create method to
	// store the information for the new user in the database. Return an error
	// if there is an issue creating the user (this will be used during dev).
	if err := u.us.Create(&user); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Fprintln(w, "User is", user)
}

// This is the struct to hold the information submitted using the Signup form on
// the /signup webpage
type SignupForm struct {
	// Use struct tags so gorilla/schema package knows about the input fields
	// in the Signup form.
	Email    string `schema:"email"`
	Name     string `schema:"name"`
	Password string `schema:"password"`
}
