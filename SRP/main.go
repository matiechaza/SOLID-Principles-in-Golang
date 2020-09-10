package SRP

import "golang.org/x/crypto/bcrypt"

type User struct {
	password string
}

/**
 * Bad: because User knows how to encrypt the password.
 */
func (u *User) SetPassword(password string) {
	u.password = u.HashPasword(password)
}

func (u *User) HashPasword(password string) string {
	hash, _ := bcrypt.GenerateFromPassword([]byte(password), 15)
	return string(hash)
}

/**
 * Good: because another structure encrypts the password.
 */
func (u *User) SetPassword(password string, hasher PasswordHasher) {
	u.password = hasher.hashed(password)
}

type PasswordHasher struct {}

func (h PasswordHasher) hashed(password string) string {
	hash, _ := bcrypt.GenerateFromPassword([]byte(password), 15)
	return string(hash)
}