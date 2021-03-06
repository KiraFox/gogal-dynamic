package models

import (
	"errors"

	"github.com/KiraFox/gogal-dynamic/hash"
	"github.com/KiraFox/gogal-dynamic/rand"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres" //loads postgres driver
	"golang.org/x/crypto/bcrypt"                 // used for hashing passwords
)

var (
	// ErrNotFound is returned when a resource cannot be found in the database.
	ErrNotFound = errors.New("models: resource not found in database")
	// ErrInvalidId is returned when an invalid ID is provided to a method like
	// Delete.
	ErrInvalidID = errors.New("models: ID provided was invalid")
	// ErrInvalidPassword is returned when an invalid password is used when
	// attempting to authenticate a user.
	ErrInvalidPassword = errors.New("models: incorrect password provided")
	userPwPepper       = "secret-random-string"
)

// Key for our HMAC for hashing remember tokens. To be used in dev only
const hmacSecretKey = "secret-hmac-key"

// This is the struct for the User model and contains what attributes(fields) we
// want to know(store) for each user of the application in the users table.
type User struct {
	gorm.Model
	Name         string
	Email        string `gorm:"not null;unique_index"`
	Password     string `gorm:"-"`                     // Raw user password, not saved
	PasswordHash string `gorm:"not null"`              // Hashed user password, saved
	Remember     string `gorm:"-"`                     // Raw remember token, not saved
	RememberHash string `gorm:"not null;unique_index"` // Hashed remember token, saved
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
	// This initializes our HMAC to be used for hashing remember tokens as needed.
	hmac := hash.NewHMAC(hmacSecretKey)

	return &UserService{
		db:   db,
		hmac: hmac,
	}, nil
}

// Abstraction layer for the users database. Provide methods for querying,
// creating, and updating users.
type UserService struct {
	db *gorm.DB // The gorm.DB object we want to interact with instead of acting
	// on the object directly throughout all our code.
	hmac hash.HMAC
}

// This method closes the UserService database connection
func (us *UserService) Close() error {
	return us.db.Close()
}

// This method will attempt to automatically migrate the users table.
func (us *UserService) AutoMigrate() error {
	if err := us.db.AutoMigrate((&User{})).Error; err != nil {
		return err
	}
	return nil
}

// This method drops (deletes) the user table then rebuilds it. This is for use
// during development.
func (us *UserService) DestructiveReset() error {
	err := us.db.DropTableIfExists(&User{}).Error
	if err != nil {
		return err
	}
	return us.AutoMigrate()
}

// This method will create the provided user and backfill data.
// Input a pointer to the User model object, run bcrypt hashing on the Password
// field from the model and check for errors.  If there are no errors, then save
// the hashed password created in model's PasswordHash field and change the
// password field to an empty string to clear the original one input. Next we
// check if Remember field is empty, and if so, set a remember token value. Then
// hash the remember token and save it.
// Run the gorm.DB Create method on the modified user model to save the User
// data to the database, and then return an error or nil if it runs successfully.
func (us *UserService) Create(user *User) error {
	pwBytes := []byte(user.Password + userPwPepper)
	hashedBytes, err := bcrypt.GenerateFromPassword(
		pwBytes, bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user.PasswordHash = string(hashedBytes)
	user.Password = ""

	if user.Remember == "" {
		token, err := rand.RememberToken()
		if err != nil {
			return err
		}
		user.Remember = token
	}

	user.RememberHash = us.hmac.Hash(user.Remember)

	return us.db.Create(user).Error
}

// This method will update the provided user with all the data in the provided
// User model object.
func (us *UserService) Update(user *User) error {
	if user.Remember != "" {
		user.RememberHash = us.hmac.Hash(user.Remember)
	}

	return us.db.Save(user).Error
}

// This method will delete the user with provided ID.
// Check if ID is == 0 to give error instead of deleting all users.
func (us *UserService) Delete(id uint) error {
	if id == 0 {
		return ErrInvalidID
	}
	user := User{Model: gorm.Model{ID: id}}
	return us.db.Delete(&user).Error
}

// This method queries database by the user ID given.
// Create intial query using GORM and save it to a variable to be used in the
// function "first" to look through the provided database and run the First
// method for us, and return the information to the User model or return errors.
func (us *UserService) ByID(id uint) (*User, error) {
	var user User
	db := us.db.Where("id = ?", id)
	err := first(db, &user)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

// This method looks up a user with the given email address.
// If the user is found, we will return a nil error.
// If the user is not found, we will return ErrNotFound.
// If there is another error, we will return an error with more information
// about what went wrong. May not be an error generated by the models package.
func (us *UserService) ByEmail(email string) (*User, error) {
	var user User
	db := us.db.Where("email = ?", email)
	err := first(db, &user)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

// This method looks up a user with the provided remember token and returns that
// user.  This method will handle hashing the provided token and searching the
// database for the matching hash. GORM uses snake case of all fields it creates
// in the database, so remember_hash = RememberHash field.
func (us *UserService) ByRemember(token string) (*User, error) {
	var user User
	rememberHash := us.hmac.Hash(token)
	err := first(us.db.Where("remember_hash = ?", rememberHash), &user)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

// This method can be used to authenticate a user with the provided email
// address and password.
// If email provided is invalid, this will return nil, ErrNotFound.
// If password provided is invalid, this will return nil, ErrInvalidPassword.
// If email and password are both valid, this will return user, nil.
// If another error is encountered, this will return nil, error.
func (us *UserService) Authenticate(email, password string) (*User, error) {
	// ByEmail method returns ErrNotFound, so we just check if that error is present.
	foundUser, err := us.ByEmail(email)
	if err != nil {
		return nil, err
	}
	// This function accepts 2 parameters - hashed password and cleartext password
	// provided. Both need to be byte slices. Also need to add pepper to cleartext
	// password as that was added before the hashing so it's needed to validate
	// the cleartexr password.
	err = bcrypt.CompareHashAndPassword(
		[]byte(foundUser.PasswordHash), []byte(password+userPwPepper))
	// Check the error return by the password comparison and either return user,
	// invalid password error, or a different error.
	switch err {
	case nil:
		return foundUser, nil
	case bcrypt.ErrMismatchedHashAndPassword:
		return nil, ErrInvalidPassword
	default:
		return nil, err
	}
}

// This function is used to query the database given and return the First record
// found matching, and stores the information in a given destination or returns
// errors if unable to complete the query.
func first(db *gorm.DB, dst interface{}) error {
	err := db.First(dst).Error
	if err == gorm.ErrRecordNotFound {
		return ErrNotFound
	}
	return err
}
