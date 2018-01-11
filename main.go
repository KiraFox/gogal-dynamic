package main

import (
	"fmt"
	"net/http"
)

/*
This is the function that we use to process incoming web request: routing
using If/Else statements

-- r.URL.Path == "/"  : First we get the URL from the Request object (r), which
returns a *url.URL. This struct has a field on it named Path that returns the
path of a URL. For example, if the URL was http://lenslocked.com/docs/abc then
the Path would be /docs/abc. The only exception here is that an empty path is
always set to "/"" , or the root path (home page).
-- Once we have the path, we use it to determine what page to render (show the
user). When the user is visiting the root path (/) we return our Welcome page.
If the user is visiting our contact page (/contact) we return a page with
information on how to contact.
-- If neither of these criteria are met, we write the 404 HTTP status code and
then write an error message to the response writer (w). The StatusNotFound vari-
able is really just a constant representing the HTTP status code 404, and the
WriteHeader method is one way to write HTTP status codes in Go.
--  http.StatusNotFound is really just a constant that represents the 404 status
code. This isn’t actually necessary, and you could replace it with 404 in your
code, but constants like StatusNotFound are often exported by packages to make
code easier to read and maintain.
--  This is very hard to maintain with multiple pages. Instead break pages into
their own function and use a router to call the function for the page.
*/
func handlerFunc(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	if r.URL.Path == "/" {
		fmt.Fprint(w, "<h1>Welcome to my awesome site!</h1>")
	} else if r.URL.Path == "/contact" {
		fmt.Fprint(w, "To get in touch, please send an email "+
			"to <a href=\"mailto:support@lenslocked.com\">"+
			"support@lenslocked.com</a>.")
	} else {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprint(w, "<h1> We could not find the page you "+
			"were looking for.</h1>"+
			"<p>Please email us if you keep being sent to an "+
			"invalid page.</p>")
	}
}

/*
This is the function that will be run to start up your application, so it needs
to call any other code that you want to run:

-- We set our handlerFunc as the function we want to use to handle web requests
going to our server with the path "/". This covers all paths that the user
might try to visit on our website. Ex: could also visit http://localhost:3000/
some-other-path and it would also be processed by handlerFunc
-- We call http.ListenAndServe(":3000", nil), which starts up a web server
listening on port 3000 using the default http handlers.
-- The port comes from the last part of the URL. You don’t have to include a
port because the browser will use a default port automatically, but you could
type it explicitly if you wanted when visiting websites.
*/

func main() {
	http.HandleFunc("/", handlerFunc)
	http.ListenAndServe(":3000", nil)
}
