package main

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/gorilla/mux"
)

var homeTemplate *template.Template
var contactTemplate *template.Template

func main() {

	var err error
	// Alter both ParseFiles for home and contact templates so both Views can
	// use (access) the new footer template layout.
	homeTemplate, err = template.ParseFiles(
		"views/home.gohtml",
		"views/layouts/footer.gohtml")
	if err != nil {
		panic(err)
	}

	contactTemplate, err = template.ParseFiles(
		"views/contact.gohtml",
		"views/layouts/footer.gohtml")
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

	if err := homeTemplate.Execute(w, nil); err != nil {
		panic(err)
	}
}

// This is the function the router calls when a user visits the "contact" page.
func contact(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")

	if err := contactTemplate.Execute(w, nil); err != nil {
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
