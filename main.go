package main

import (
	"fmt"
	"net/http"

	"github.com/KiraFox/gogal-dynamic/controllers"
	"github.com/KiraFox/gogal-dynamic/models"

	"github.com/gorilla/mux"
)

// Database information. Only use this during development. Do not commit to git
// with information.
const (
	host     = "localhost"
	port     = 5432
	user     = "user"
	password = "password"
	dbname   = "database"
)

func main() {
	// Create a database connection string
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	// Use connection string to create Model services and start it and defer
	// close until application is done.
	us, err := models.NewUserService(psqlInfo)
	if err != nil {
		panic(err)
	}
	defer us.Close()
	us.AutoMigrate()

	staticC := controllers.NewStatic()
	usersC := controllers.NewUsers(us)

	var nf http.Handler
	nf = http.HandlerFunc(notFound)

	r := mux.NewRouter()

	r.Handle("/", staticC.Home).Methods("GET")
	r.Handle("/contact", staticC.Contact).Methods("GET")
	r.HandleFunc("/signup", usersC.New).Methods("GET")
	r.HandleFunc("/signup", usersC.Create).Methods("POST")
	// Use Handle instead of HandleFunc to render the login webpage like a static
	// page
	r.Handle("/login", usersC.LoginView).Methods("GET")
	r.HandleFunc("/login", usersC.Login).Methods("POST")
	r.HandleFunc("/cookietest", usersC.CookieTest).Methods("GET")
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
