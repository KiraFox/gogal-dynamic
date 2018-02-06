package rand

import (
	"crypto/rand"
	"encoding/base64"
)

const RememberTokenBytes = 32

// This function will help us generate n random bytes slice, or will return an
// error if there was one. This uses the crypto/rand package.
func Bytes(n int) ([]byte, error) {
	// Creates a byte slice of size n
	b := make([]byte, n)
	_, err := rand.Read(b)
	if err != nil {
		return nil, err
	}
	return b, nil
}

// This function will generate a random byte slice of size nBytes using our
// Bytes function then return a string that is the base64 URL encoded version of
// the byte slice.
func String(nBytes int) (string, error) {
	b, err := Bytes(nBytes)
	if err != nil {
		return "", err
	}
	return base64.URLEncoding.EncodeToString(b), nil
}

// This function is a helper function designed to generate remember tokens of a
// predetermined (const) byte size.
func RememberToken() (string, error) {
	return String(RememberTokenBytes)
}
