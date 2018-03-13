package controllers

import (
	"fmt"
	"net/http"

	"github.com/KiraFox/gogal-dynamic/models"
	"github.com/KiraFox/gogal-dynamic/rand"
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

	// Sign the newly created user in using the signIn method
	// Temporarily renders error message for debugging and redirects to the
	// cookie test page to make sure the signIn worked and cookie was set
	err := u.signIn(w, &user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	http.Redirect(w, r, "/cookietest", http.StatusFound)
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

	// Sign the newly created user in using the signIn method
	// Temporarily renders error message for debugging and redirects to the
	// cookie test page to make sure the signIn worked and cookie was set
	err = u.signIn(w, user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	http.Redirect(w, r, "/cookietest", http.StatusFound)
}

// This method is used to display cookies set on the current user
func (u *Users) CookieTest(w http.ResponseWriter, r *http.Request) {
	// Use the Name field of the cookie you want to see the values of and check
	// for errors to see if the cookie was located.
	cookie, err := r.Cookie("remember_token")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	// Use this cookie's value to search for a user in our database. We are using
	// the ByRemember method search that we created for the UserService
	user, err := u.us.ByRemember(cookie.Value)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	fmt.Fprintln(w, user)
}

// This method is used to sign the given user in via cookies
func (u *Users) signIn(w http.ResponseWriter, user *models.User) error {
	// Check if a raw remember token is set: if not set - generate a new remember
	// token then update the user; otherwise use the remember token that is
	// already set for the user
	if user.Remember == "" {
		token, err := rand.RememberToken()
		if err != nil {
			return err
		}
		user.Remember = token
		err = u.us.Update(user)
		if err != nil {
			return err
		}
	}
	// Use the remember token to create a cookie and set it using the http
	// package along with the response writer
	cookie := http.Cookie{
		Name:  "remember_token",
		Value: user.Remember,
	}
	http.SetCookie(w, &cookie)

	return nil
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
