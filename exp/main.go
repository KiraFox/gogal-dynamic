package main

import (
	"html/template"
	"os"
)

func main() {
	// ParseFiles will open up the template file and attempt to validate it. If
	// everything goes well, you will receive a *Template and a nil error,
	// otherwise you will receive a nil template and an error.
	t, err := template.ParseFiles("hello.gohtml")

	// After calling ParseFiles we check to see if an error was returned. If an
	// error is present we panic. Otherwise we continue on with our program
	// knowing we have a valid template referenced by the t variable.
	if err != nil {
		panic(err)
	}

	// Create the variable (data) as an anonymous struct with a field (Name) of
	// type (string).
	// When we initialize (data) we set the Name field to “John Smith"
	data := struct {
		Name string
	}{"John Smith"}

	// Execute the template parsed earlier (hello.gohtml), passing in two
	// arguments:
	// 1. Where we want to write the template output
	//	+  In this case: we want to write the template output to Stdout, which
	//	is your terminal window.
	// 2. The data to be used when executing the template
	//	+  In this case: in the hello.gohtml file, we use the line {{.Name}}
	//	to render a name dynamically, so we need to pass in a data structure
	//	with a Name field. The variable (data) that we created has this
	//	field set to “John Smith”, so when we provide it to the template we
	//	should see the name “John Smith” in the final output.

	err = t.Execute(os.Stdout, data)
	if err != nil {
		panic(err)
	}
}
