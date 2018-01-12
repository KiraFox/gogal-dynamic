package views

import (
	"html/template"
	"path/filepath" // Use for Globbing
)

// Create global variables to help construct Glob pattern
// Technically these can be constants but leave them as variables in case we
// want to test this code with different values later.
var (
	LayoutDir   string = "views/layouts/" // Specifies layout directory
	TemplateExt string = ".gohtml"        // Specifies file extension for templates
)

// Create function to use the variables & Glob function then return slice
// of templates to include in our view
func layoutFiles() []string {

	// Pass in a string to the Glob function that we create via combining
	// variables and a hard-coded * string (="views/layouts/*.gohtml").
	files, err := filepath.Glob(LayoutDir + "*" + TemplateExt)
	if err != nil {
		panic(err)
	}

	// Return all of the file paths we receive from our call to Glob
	return files
}

// Simplify funcation to use the new Glob function (layoutFiles)
func NewView(layout string, files ...string) *View {
	// Instead of hard-coding individual files to append, we pass in all the
	// files returned by layoutFiles
	files = append(files, layoutFiles()...)

	t, err := template.ParseFiles(files...)
	if err != nil {
		panic(err)
	}

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
