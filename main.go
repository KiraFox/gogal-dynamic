package main

import (
	"fmt"
	"net/http"

	"github.com/KiraFox/gogal-dynamic/controllers"

	"github.com/gorilla/mux"
)

func main() {
	// Create static controller for home/contact/etc pages
	staticC := controllers.NewStatic()
	usersC := controllers.NewUsers()

	var nf http.Handler
	nf = http.HandlerFunc(notFound)

	r := mux.NewRouter()
	// Update router to use Handle for static pages"
	r.Handle("/", staticC.Home).Methods("GET")
	r.Handle("/contact", staticC.Contact).Methods("GET")
	r.HandleFunc("/signup", usersC.New).Methods("GET")
	r.HandleFunc("/signup", usersC.Create).Methods("POST")
	r.NotFoundHandler = nf
	http.ListenAndServe(":3000", r)
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
