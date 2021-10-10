package crypto

import (
	"encoding/base64"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

func PasswordEncrypt(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(hash), err
}

func CompareHashAndPassword(hash, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
}

func SecureRandomBase64() string {
	return base64.StdEncoding.EncodeToString(uuid.New().NodeID())
}
