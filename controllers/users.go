package controllers

import (
	"net/http"

	"github.com/KiraFox/gogal-dynamic/views"
)

// This is the controller for the "users" resource
type Users struct {
	//Update controller to store a "new" user view template
	NewView *views.View
}

// This function handles the logic for rendering "new" users view template and
// returning the information to the Users controller struct
func NewUsers() *Users {
	return &Users{
		// Call the NewView function in our created views package to parse
		// which templates we want to use for the "new" users page
		NewView: views.NewView("bootstrap", "views/users/new.gohtml"),
	}
}

// This method is used to render the view stored in the NewView field of the
// Users controller struct for the Sign Up ("new" users) webpage
// GET /signup
func (u *Users) New(w http.ResponseWriter, r *http.Request) {
	// Using "u" to access the Users controller struct
	// Since NewView holds a views.View we can call the Render method from the
	// associed with the View type struct
	if err := u.NewView.Render(w, nil); err != nil {
		panic(err)
	}
}
