package lib

import "github.com/pilinux/argon2"

func CreateHash(password string, secret string) string {
	hasher, _ := argon2.CreateHash(password, secret, argon2.DefaultParams)
	return hasher
}

func HashValidator(password string, secret string, hash string) bool {
	isValid, _ := argon2.ComparePasswordAndHash(password, secret, hash)
	return isValid
}
