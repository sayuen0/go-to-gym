package utils

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

// PasswordVerifier manages password verification
type PasswordVerifier struct {
	Salt   string
	Pepper string
}

// GenerateSalt generates a new salt string for password hashing
func GenerateSalt() string {
	return NewUUIDStr()
}

// HashPassword hashes its password string with its secret-salt and salt string
func (s PasswordVerifier) HashPassword(rawPassword string) (string, error) {
	hashed, err := bcrypt.GenerateFromPassword(
		[]byte(s.getSeasonedPassword(rawPassword)), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	return string(hashed), nil
}

// ComparePassword returns an error if the hash of input password is not equal to its hashed password
// TODO: 引数どっちがどっちかわからなくなるので、型をつけるなりして対処したい
func (s PasswordVerifier) ComparePassword(
	inputPassword, hashedPassword string,
) error {
	if err := bcrypt.CompareHashAndPassword(
		[]byte(hashedPassword), []byte(s.getSeasonedPassword(inputPassword))); err != nil {
		return err
	}

	return nil
}

// getSeasonedPassword concatenates password string with its secret-salt and salt string
func (s PasswordVerifier) getSeasonedPassword(rawPassword string) string {
	return fmt.Sprintf("%s%s%s", rawPassword, s.Pepper, s.Salt)
}
