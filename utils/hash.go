package utils

import "github.com/matthewhartstonge/argon2"

var argon = argon2.DefaultConfig()

func HashPassword(password string) string {
	e, _ := argon.HashEncoded([]byte(password))
	return string(e)
}

func VerifyPassword(password string, hash string) bool {
	ok, _ := argon2.VerifyEncoded([]byte(password), []byte(hash))
	return ok
}
