package security

import (
	"github.com/alexedwards/argon2id"
)

func GeneratePasswordHash(password []byte) ([]byte, error) {
	hash, err := argon2id.CreateHash(string(password), argon2id.DefaultParams)
	return []byte(hash), err
}

func CheckPassword(password string, passwordHash string) bool {
	match, err := argon2id.ComparePasswordAndHash(password, passwordHash)
	if err != nil {
		panic(err)
	}

	return match
}
