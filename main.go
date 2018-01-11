package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

/*
Routing via Gorilla Mux:

-- Have a function for each path(webpage).
-- First we create a new router: r := mux.NewRouter()
-- Then we start assigning functions to handle different paths:
r.HandleFunc("/", home) & r.HandleFunc("/contact", contact)
-- Finally we start up our server: http.ListenAndServe(":3000", r)
	+ This time we are passing in our router (r) as the default handler for web
	requests instead of "nil" like before. This tells the ListenAndServe
	function that we want to use our own custom router.
	+ Our router will in turn handle requests long enough to figure out which
	function was assigned to that path, and then it will call that function.
-- gorilla/mux provides a simple 404 status page by default
	+ Custimize the 404 page by setting the NotFoundHandler attribute on the
	gorilla/mux.Router
	+ You need to initialize a variable as a http.Handler interface:
	var nf http.Handler
	+ Then set the variable (nf) to a http.HandlerFunc(...) with a function you
	created for the 404 webpage:
	nf = http.HandlerFunc(notFound)
	+ Then set the NotFoundHandler on the gorilla/mux.Router to your variable:
	r.NotFoundHandler = nf
*/
func main() {
	var nf http.Handler
	nf = http.HandlerFunc(notFound)

	r := mux.NewRouter()
	r.HandleFunc("/", home)
	r.HandleFunc("/contact", contact)
	r.NotFoundHandler = nf
	http.ListenAndServe(":3000", r)
}

// This is the function the router calls when a user visits the "home" page.
func home(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	fmt.Fprint(w, "<h1>Welcome to my site!</h1>")
}

// This is the function the router calls when a user visits the "contact" page.
func contact(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	fmt.Fprint(w, "To get in touch, please send an email "+
		"to <a href=\"mailto:support@lenslocked.com\">"+
		"support@lenslocked.com</a>.")
}

// This is the function for a custom 404 status page.
func notFound(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotFound)
	fmt.Fprint(w, "<h1> We could not find the page you "+
		"were looking for.</h1>"+
		"<p>Please email us if you keep being sent to an "+
		"invalid page.</p>")
}
