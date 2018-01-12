package views

import (
	"html/template"
)

// This funcation was created in order to make it easier to create Views.
// Function now takes in new parameter (layout) so we can set this on the view
// we are creating.
func NewView(layout string, files ...string) *View {
	// Add the navbar.gohtml layout to the slice of template files we are
	// parsing so it is available for rendering.
	files = append(files,
		"views/layouts/footer.gohtml",
		"views/layouts/bootstrap.gohtml",
		"views/layouts/navbar.gohtml")

	// Use the ... operator after a variable name to “unravel” the items in a slice
	// template.ParseFiles() expects strings, not a slice, so "files..." allows
	// us to take the slice of strings and treat it as a list of strings
	t, err := template.ParseFiles(files...)
	if err != nil {
		panic(err)
	}
	// This sets the Template field to whatever parsed template files we used.
	// Set Layout field to the layout that is called for by the parameter when
	// the function is called.
	return &View{
		Template: t,
		Layout:   layout,
	}
}

// Declaring our View type
// Add new field Layout to store the name of the template we want the view to
// Execute
type View struct {
	Template *template.Template
	Layout   string
}
