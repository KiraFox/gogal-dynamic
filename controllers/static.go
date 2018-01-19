package controllers

import "github.com/KiraFox/gogal-dynamic/views"

// This is the controller for individual static webpages
type Static struct {
	Home    *views.View
	Contact *views.View
}

// This function handles the logic for parsing each static page's view template
// and returning the information to the Static controller struct
func NewStatic() *Static {
	return &Static{
		Home:    views.NewView("bootstrap", "views/static/home.gohtml"),
		Contact: views.NewView("bootstrap", "views/static/contact.gohtml"),
	}
}
