package views

import (
	"html/template"
	"net/http"
	"path/filepath"
)

var (
	LayoutDir   string = "views/layouts/"
	TemplateDir string = "views/"
	TemplateExt string = ".gohtml"
)

// Function is used to put all template layout files from specific directory and
// of a certain file type into a slice.
func layoutFiles() []string {

	files, err := filepath.Glob(LayoutDir + "*" + TemplateExt)
	if err != nil {
		panic(err)
	}

	return files
}

// Function that takes in files and adds the template directory and extension
func addTempPathExt(files []string) {
	for i, f := range files {
		files[i] = TemplateDir + f + TemplateExt
	}
}

// Function is used to populate our View type with the parsed template layout
// files gathered by the layoutFiles function and the name of the layout used.
func NewView(layout string, files ...string) *View {
	// Call function to add template path and file extention to all the files
	// being used to create a View
	addTempPathExt(files)
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

// Declaring our View type to hold pointer to parsed templates and name of layout
// of the webpage
type View struct {
	Template *template.Template
	Layout   string
}

// This renders the view of the webpage and handles the logic so it can be used
// by the handler functions instead of the logic being coded in the handlers.
// The data parameter will be used in the future.
func (v *View) Render(w http.ResponseWriter, data interface{}) error {
	w.Header().Set("Content-Type", "text/html")
	return v.Template.ExecuteTemplate(w, v.Layout, data)
}

// Create new method for View type that receives an http response writer and a
// request so Views rendered can be used by router without a separate handler
// function if needed. ServeHTTP(ResponseWriter,*Request) = http.Handler object
func (v *View) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if err := v.Render(w, nil); err != nil {
		panic(err)
	}
}
