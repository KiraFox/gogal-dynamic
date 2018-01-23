package models

import (
	"errors"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres" //loads postgres driver
)

var (
	// ErrNotFound is returned when a resource cannot be found in the database.
	ErrNotFound = errors.New("models: resource not found in database")
)

// This is the struct for the User model and contains what attributes(fields) we
// want to know(store) for each user of the application in the users table.
type User struct {
	gorm.Model
	Name  string
	Email string `gorm:"not null;unique_index"`
}

// Abstraction layer for the users database. Provide methods for querying,
// creating, and updating users.
type UserService struct {
	db *gorm.DB // The gorm.DB object we want to interact with instead of acting
	// on the object directly throughout all our code.
}

// This function opens a connection to a database (connectionInfo) and returns
// a gorm.DB object to our UserService to use. It also logs the SQL run. Do not
// close the database in this function as it would close before it returns the
// object to the UserService.
func NewUserService(connectionInfo string) (*UserService, error) {
	// End user passes a string defining how to connect to the database and we
	// want to use the postgres dialect.
	db, err := gorm.Open("postgres", connectionInfo)
	if err != nil {
		return nil, err
	}

	db.LogMode(true)
	return &UserService{
		db: db,
	}, nil
}

// This method closes the UserService database connection
func (us *UserService) Close() error {
	return us.db.Close()
}

// This method queries database by user ID. ByID will look up a user with the
// provided ID. If the user is found, we return nil error and information to
// User model. If user is not found	we return the custom error variable. If
// there is a different error, we return whatever error was generated.
func (us *UserService) ByID(id uint) (*User, error) {
	var user User
	err := us.db.Where("id = ?", id).First(&user).Error
	switch err {
	case nil:
		return &user, nil
	case gorm.ErrRecordNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

// This method drops (deletes) the user table then rebuilds it. This is for use
// during development.
func (us *UserService) DestructiveReset() {
	us.db.DropTableIfExists(&User{})
	us.db.AutoMigrate(&User{})
}
