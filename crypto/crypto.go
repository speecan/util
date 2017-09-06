package crypto

import (
	"github.com/labstack/gommon/random"

	"golang.org/x/crypto/bcrypt"
)

// HashPassword returns password hash
func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

// CheckPasswordHash checks validation
func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

// Random returns random string with random.Alphanumeric
func Random(n int) string {
	return random.String(uint8(n), random.Alphanumeric)
}
