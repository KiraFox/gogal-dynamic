package views

import (
	"html/template"
)

// This funcation was created in order to make it easier to create Views.
// This function is going to handle appending common template files to the list
// of files provided, parse the template files, then construct the *View (pointer
// to our View struct type) and return it for use.
// Function takes in any number of strings (...string) as its argument then it
// merges them into a string slice ([]slice) named "files" for us to use
func NewView(files ...string) *View {
	files = append(files, "views/layouts/footer.gohtml")

	// Use the ... operator after a variable name to “unravel” the items in a slice
	// template.ParseFiles() expects strings, not a slice, so "files..." allows
	// us to take the slice of strings and treat it as a list of strings
	t, err := template.ParseFiles(files...)
	if err != nil {
		panic(err)
	}
	// This sets the Template field to whatever parsed template files we used
	return &View{
		Template: t,
	}
}

// Declaring our View type
// This is a struct that contains one attribute - a pointer to a template.Template
// which will eventually point to our compiled template.
// This will change over time but it needs to store the parsed template we want
// to Execute.
type View struct {
	Template *template.Template
}
