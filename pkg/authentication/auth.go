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
