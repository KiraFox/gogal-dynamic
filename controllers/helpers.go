package controllers

import (
	"net/http"

	"github.com/gorilla/schema"
)

// This will be used anytime we need to parse a form submitted and then decode it
// using the gorilla/schema. This is to help keep code DRY.
// The parameter (r) is where we are getting the form from, and the parameter
// (dst) is the destination we want the final information stored. (dst) is an
// empty interface type so we can have our destination be any type we need.
//
func parseForm(r *http.Request, dst interface{}) error {
	// Parse the form submitted that is part of the http.Request (r) and return
	// any errors encountered.
	if err := r.ParseForm(); err != nil {
		return err
	}

	// Intialize gorilla/schema decoder to use on the parsed form.
	dec := schema.NewDecoder()

	// Decode the parsed data (r.PostForm) then store it in the destination (dst)
	// and return any errors encountered.
	if err := dec.Decode(dst, r.PostForm); err != nil {
		return err
	}
	// Returns nil errors if no errors were encountered when running the previous
	// code.
	return nil
}
