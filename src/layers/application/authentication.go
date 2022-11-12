package application

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
	"golang.org/x/crypto/pbkdf2"
)

var innerSalt = []byte{231, 15, 88, 230, 39, 206, 151, 15}

func hashing(password string, salt []byte, iter, keyLen int) string {
	passwordInBytes := []byte(password)

	now := pbkdf2.Key(passwordInBytes, salt, iter, keyLen, sha256.New)
	return hex.EncodeToString(now)
}

func getSalt() (b []byte, err error) {
	bigInt, err := rand.Prime(rand.Reader, 256)
	if err != nil {
		return
	}

	b = bigInt.Bytes()
	return
}
