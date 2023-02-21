package models

import (
	"golang.org/x/crypto/bcrypt"
	"strings"
	"time"
)

// TODO: create user create, update request
// TODO: write validate tag

// ---------------------------------------------------------------------------------------------------------------------
// user get response

type User struct {
	UserID         string    `json:"user_id"`
	Name           string    `json:"name"`
	Email          string    `json:"email"`
	Password       string    `json:"password"`
	HashedPassword string    `json:"-"`
	Gender         *string   `json:"gender,omitempty"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
	LoginDate      time.Time `json:"login_date"`
}

func (u *User) HashPassword() error {
	hashed, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	u.HashedPassword = string(hashed)
	return nil
}

func (u *User) ComparePasswords(password string) error {
	if err := bcrypt.CompareHashAndPassword([]byte(u.HashedPassword), []byte(password)); err != nil {
		return err
	}
	return nil
}

func (u *User) SanitizePassword() {
	u.Password = ""
}

func (u *User) PrepareCreate() error {
	u.Email = strings.ToLower(strings.TrimSpace(u.Email))
	u.Password = strings.TrimSpace(u.Password)

	if err := u.HashPassword(); err != nil {
		return err
	}
	return nil
}

func (u *User) PrepareUpdate() error {
	u.Email = strings.ToLower(strings.TrimSpace(u.Email))
	return nil
}

type UserWithToken struct {
	User  *User  `json:"user"`
	Token string `json:"token"`
}
