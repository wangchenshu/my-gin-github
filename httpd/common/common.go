package common

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

// HashPassword hash password
func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

// CheckPasswordHash check password hash
func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

// MyRecover - recover panic
func MyRecover() {
	if e := recover(); e != nil {
		fmt.Println(e)
	}
}
