package main

import (
	"fmt"
	"net/http"
)

/*
This is the function that we use to process incoming web request

-- Every time someone visits your website, the code in handlerFunc(...) gets
run and determines what to return to the visitor
-- We will have different handlers for when users visit different pages on our
application
-- All handlers take the same two arguments. An http.ResponseWriter declared as
w in our current code, and a pointer to an http.Request, declared as r in our
current code.
-- http.ResponseWriter : This is a structure that allows us to modify the
response that we want to send to whoever visited our website. Default
implements the Write() method that allows us to write to the response body. Has
methods that help us set headers when we need to
-- *http.Request : This is a structure used to access data from the web
request. For example, we might use this to get the users email address and
password after they sign up for our web application.
*/
func handlerFunc(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "<h1>Welcome to my awesome site!</h1>")
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
