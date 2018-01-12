package main

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/gorilla/mux"
)

// Using MVC : Create first View and Template
// Create global variable to store the parsed template for home page instead of
// using Fprint in the "home" func. Will be cleaned up later.
var homeTemplate *template.Template

func main() {

	var err error
	// Make sure global variable is assigned to the parsed template
	// Template paths are relative. Note that the file paths used to access them
	// are relative to wherever you run your code.
	// Parsing files can cause issues (return error) so check for errors
	homeTemplate, err = template.ParseFiles("views/home.gohtml")
	if err != nil {
		panic(err)
	}

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
	// Update function to use (Execute) new template (homeTemplate)
	// We write the results to our Response Writer (w), and don't give any data
	// to the template to use (data = nil) as the template doesn't currently
	// use any dynamic data
	// Executing templates can cause issues (return error) so check for errors
	if err := homeTemplate.Execute(w, nil); err != nil {
		panic(err)
	}
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
