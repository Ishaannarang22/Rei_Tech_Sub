package authentication

import (
	"golang.org/x/crypto/bcrypt"
)

// Verify the password
func CheckPass(hash, pass string) (bool, error) {
	password := []byte(pass)
	hashedPassword := []byte(hash)
	err := bcrypt.CompareHashAndPassword(hashedPassword, password)

	if err != nil {
		return false, err
	} else {
		return true, nil
	}
}

// Encrypt the password using bcrypt
func EncryptPassword(password string) (string, error) {

	hash, err := bcrypt.GenerateFromPassword([]byte(password), 10)

	if err != nil {
		return "", err
	} else {
		return string(hash), nil
	}

}
