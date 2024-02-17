package validators

import (
	"golang.org/x/crypto/bcrypt"
)


func Hash(password string) []byte {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return []byte(err.Error())
	}
	return hash
}