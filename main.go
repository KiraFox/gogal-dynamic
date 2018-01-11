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
*/
func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", home)
	r.HandleFunc("/contact", contact)
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
