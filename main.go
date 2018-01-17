package main

import (
	"fmt"
	"net/http"

	"github.com/KiraFox/gogal-dynamic/controllers"
	"github.com/KiraFox/gogal-dynamic/views"

	"github.com/gorilla/mux"
)

var homeView *views.View
var contactView *views.View

func main() {

	homeView = views.NewView("bootstrap", "views/home.gohtml")
	contactView = views.NewView("bootstrap", "views/contact.gohtml")
	// Create Users Controller variable and set it to NewUser function
	usersC := controllers.NewUsers()

	var nf http.Handler
	nf = http.HandlerFunc(notFound)

	r := mux.NewRouter()
	r.HandleFunc("/", home)
	r.HandleFunc("/contact", contact)
	// Update signup webpage to use the User Controller var and run the method
	// New for "new" users.  The method will handle any web requests for the
	// signup page now.  The router can use the method because it uses the
	// arguments of type ResponseWrite and Request.
	r.HandleFunc("/signup", usersC.New)
	r.NotFoundHandler = nf
	http.ListenAndServe(":3000", r)
}

// This is the function the router calls when a user visits the "home" page.
func home(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	must(homeView.Render(w, nil))
}

// This is the function the router calls when a user visits the "contact" page.
func contact(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	must(contactView.Render(w, nil))
}

// Helper function to check for errors and panic if one is found
func must(err error) {
	if err != nil {
		panic(err)
	}
}

// This is the function for a custom 404 status page.
func notFound(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotFound)
	fmt.Fprint(w, "<h1> We could not find the page you "+
		"were looking for.</h1>"+
		"<p>Please email us if you keep being sent to an "+
		"invalid page.</p>")
}
