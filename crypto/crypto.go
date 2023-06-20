package crypto

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"regexp"

	"github.com/labstack/gommon/random"

	"golang.org/x/crypto/bcrypt"
)

var (
	randomAlphanumericWithoutConfusable string
)

func init() {
	randomAlphanumericWithoutConfusable = avoidConfusableCharactersSet()
}

func avoidConfusableCharactersSet() string {
	confusable := "1lI0Oo8B3Evu"
	re := regexp.MustCompile(`[` + confusable + `]`)
	return re.ReplaceAllString(random.Alphanumeric, "")
}

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

// RandomWithoutConfusable returns random string with avoiding confusable characters
func RandomWithoutConfusable(n int) string {
	return random.String(uint8(n), randomAlphanumericWithoutConfusable)
}

// EncryptByGCM aes-gcm
// key should be either 16, 24 or 32 bytes
func EncryptByGCM(key []byte, plainText string) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}
	nonce := make([]byte, gcm.NonceSize()) // Unique nonce is required(NonceSize 12byte)
	_, err = rand.Read(nonce)
	if err != nil {
		return nil, err
	}
	cipherText := gcm.Seal(nil, nonce, []byte(plainText), nil)
	cipherText = append(nonce, cipherText...)
	return cipherText, nil
}

// DecryptByGCM aes-gcm
// key should be either 16, 24 or 32 bytes
func DecryptByGCM(key []byte, cipherText []byte) (string, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return "", err
	}
	nonce := cipherText[:gcm.NonceSize()]
	plainByte, err := gcm.Open(nil, nonce, cipherText[gcm.NonceSize():], nil)
	if err != nil {
		return "", err
	}
	return string(plainByte), nil
}

// DummyBinary returns dummy data
func DummyBinary(length uint64) []byte {
	b := make([]byte, 0)
	for uint64(len(b)) < length {
		str := Random(128)
		b = append(b, []byte(str)...)
	}
	return b
}
