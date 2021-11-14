package utils

import (
	"crypto/rand"
	"math/big"

	"golang.org/x/crypto/bcrypt"
)

const saltLen = 40

func GenerateSalt() []byte {
	const letters = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
	ret := make([]byte, saltLen)
	for i := 0; i < saltLen; i++ {
		num, _ := rand.Int(rand.Reader, big.NewInt(int64(len(letters))))
		ret[i] = letters[num.Int64()]
	}
	return ret
}

func CalculatePasswordWithSalt(originalPassword string, salt []byte) []byte {
	hashed := []byte(originalPassword)
	hashed = append(hashed, salt...)
	res, _ := bcrypt.GenerateFromPassword(hashed, bcrypt.DefaultCost)
	return res
}

// `storedPassword` should equal to `bcrypt(sha512(password) + salt)`, where sha512(password) is `hashedPassword`
// and is sent from clients.
func VerifyPassword(hashedPassword string, salt string, storedPassword string) error {
	hashed := []byte(hashedPassword)
	hashed = append(hashed, []byte(salt)...)
	return bcrypt.CompareHashAndPassword([]byte(storedPassword), hashed)
}
