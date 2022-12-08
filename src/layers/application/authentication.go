package application

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
	"golang.org/x/crypto/pbkdf2"
)

var innerSalt = []byte{231, 15, 88, 230, 39, 206, 151, 15}

const (
	numberOfHashing = 8
	keyLength       = 32
)

func hashing(password string, salt []byte) string {
	passwordInBytes := []byte(password)

	salt = append(salt, innerSalt...)

	now := pbkdf2.Key(passwordInBytes, salt, numberOfHashing, keyLength, sha256.New)
	return hex.EncodeToString(now)
}

func generateSalt() (b []byte, err error) {
	bigInt, err := rand.Prime(rand.Reader, 256)
	if err != nil {
		return
	}

	b = bigInt.Bytes()
	return
}

func generateCookieHash() {

}
