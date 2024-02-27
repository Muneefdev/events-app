package utails

import (
	"golang.org/x/crypto/bcrypt"
)

func Hash(password string) string {
	bytes, _ := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes)
}

func ComparePasswords(hashedPassword, password string) bool {
  err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
  return err == nil
}
