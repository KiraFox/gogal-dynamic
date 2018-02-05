package controllers

import (
	"fmt"
	"net/http"

	"github.com/KiraFox/gogal-dynamic/models"
	"github.com/KiraFox/gogal-dynamic/views"
)

// This function handles the logic for parsing user templates and then returns
// the information to the Users controller struct. It also takes a user service
// and assigns it to the same struct to be used by its methods.
func NewUsers(us *models.UserService) *Users {
	return &Users{
		NewView:   views.NewView("bootstrap", "users/new"),
		LoginView: views.NewView("bootstrap", "users/login"),
		us:        us,
	}
}

// This is the controller for the "users" resource
type Users struct {
	// Store a "new" user view template
	NewView *views.View
	// Store a login user view template
	LoginView *views.View
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

// This method is used to process the Login form when a user tries to login to
// an existing user account.
// POST /login
func (u *Users) Login(w http.ResponseWriter, r *http.Request) {
	form := LoginForm{}
	if err := parseForm(r, &form); err != nil {
		panic(err)
	}
	// Call the Authenticate method and provide email and password from the login
	// form and print out information based on if or what error is encountered
	// when verifying the submitted information
	user, err := u.us.Authenticate(form.Email, form.Password)
	if err != nil {
		switch err {
		case models.ErrNotFound:
			fmt.Fprintln(w, "Invalid email address.")
		case models.ErrInvalidPassword:
			fmt.Fprintln(w, "Invalid Password provided.")
		case nil:
			fmt.Fprintln(w, user)
		default:
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		return
	}
	// Create cookie with the key "email" and value of the authenticated user's
	// email address. Then save (SetCookie) the cookie.
	cookie := http.Cookie{
		Name:  "email",
		Value: user.Email,
	}
	http.SetCookie(w, &cookie)
	fmt.Fprintln(w, user)
}

// This method is used to display cookies set on the current user
func (u *Users) CookieTest(w http.ResponseWriter, r *http.Request) {
	// Use the Name field of the cookie you want to see the values of and check
	// for errors to see if the cookie was located then print out the values we
	// are checking.
	cookie, err := r.Cookie("email")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	fmt.Fprintln(w, "Email is:", cookie.Value)
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

// This is the struct to hold the information submitted using the Login form on
// the /login webpage
type LoginForm struct {
	// Use struct tags so gorilla/schema package knows about the input fields
	// in the Login form.
	Email    string `schema:"email"`
	Password string `schema:"password"`
}
