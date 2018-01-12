package main

import (
	"fmt"
	"net/http"
	// Import the views package to use the new view type created
	"github.com/KiraFox/gogal-dynamic/views"

	"github.com/gorilla/mux"
)

// Update global variables to be a pointer to a View (*View) instead of template
// and change names to reflect what they are now (views instead of templates)
var homeView *views.View
var contactView *views.View

func main() {
	// Update to use the NewView function we created to parse the corresponding
	// templates and create our views.
	homeView = views.NewView("views/home.gohtml")
	contactView = views.NewView("views/contact.gohtml")

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
	// Update to use the view variables instead of previous template variables.
	// Access the Template field of each view object (View struct) instead of
	// directly referencing the template.
	if err := homeView.Template.Execute(w, nil); err != nil {
		panic(err)
	}
}

// This is the function the router calls when a user visits the "contact" page.
func contact(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	// Update to use the view variables instead of previous template variables.
	// Access the Template field of each view object (View struct) instead of
	// directly referencing the template.
	if err := contactView.Template.Execute(w, nil); err != nil {
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
