package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "username"
	password = "your-password"
	dbname   = "appname_dev"
)

func main() {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, err := gorm.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	// Enable logging to track what SQL statements are being run behind the scene.
	// Will only see output when SQL statements are actually run
	db.LogMode(true)

	// Creates the table in the SQL database per the fields and tags set in	the
	// User struct.
	db.AutoMigrate(&User{})

	// Gets data reuired for adding a record in the SQL table, then puts the data
	// in an instance of the table's struct.  Run the Create method on the instance
	// to create the new record in the table and check for errors.
	name, email := getInfo()
	u := &User{
		Name:  name,
		Email: email,
	}
	if err = db.Create(u).Error; err != nil {
		panic(err)
	}

	fmt.Printf("%+v\n", u)

}

// This sets the names of the fields, type and contraints in the 'users' SQL table
type User struct {
	gorm.Model
	Name  string
	Email string `gorm:"not null;unique_index"`
}

// This gets the information from a user regarding their name and email
func getInfo() (name, email string) {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("What is your name?")
	name, _ = reader.ReadString('\n')
	name = strings.TrimSpace(name)
	fmt.Println("What is your email?")
	email, _ = reader.ReadString('\n')
	email = strings.TrimSpace(email)
	return name, email
}
