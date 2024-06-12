package password

import "golang.org/x/crypto/bcrypt"

var customCost = 12

func HashPassword(password string) (string, error) {
	hashedPass, err := bcrypt.GenerateFromPassword([]byte(password), customCost)
	return string(hashedPass), err
}

func IsValidPassword(encpw, password string) bool {
	return bcrypt.CompareHashAndPassword([]byte(encpw), []byte(password)) == nil
}
