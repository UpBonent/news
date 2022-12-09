package application

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
	"golang.org/x/crypto/pbkdf2"
)

const (
	salt        = 256
	cookieValue = 128
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

func generate(length int) (b []byte, err error) {
	bigInt, err := rand.Prime(rand.Reader, length)
	if err != nil {
		return
	}

	b = bigInt.Bytes()
	return
}
