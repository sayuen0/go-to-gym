package models

import (
	"github.com/sayuen0/go-to-gym/internal/models/db"
	"golang.org/x/crypto/bcrypt"
	"strings"
	"time"
)

// ---------------------------------------------------------------------------------------------------------------------
// user create request

type UserCreateRequest struct {
	Name           string `json:"name"`
	Email          string `json:"email"`
	Password       string `json:"password"`
	HashedPassword string `json:"-"`
	// TODO: add birthday
}

func (u *UserCreateRequest) PrepareForCreate() error {
	u.Email = strings.ToLower(strings.TrimSpace(u.Email))
	u.Password = strings.TrimSpace(u.Password)

	if err := u.HashPassword(); err != nil {
		return err
	}
	return nil
}

func (u *UserCreateRequest) HashPassword() error {
	hashed, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	u.HashedPassword = string(hashed)
	return nil
}

// ---------------------------------------------------------------------------------------------------------------------
// user update request

type UserUpdateRequest struct {
	UserID         string  `json:"user_id"`
	Name           string  `json:"name"`
	Email          string  `json:"email"`
	Password       string  `json:"password"`
	HashedPassword string  `json:"-"`
	Gender         *string `json:"gender,omitempty"`
}

// TODO: write validate tag

// ---------------------------------------------------------------------------------------------------------------------
// user get response

type User struct {
	UserID    string    `json:"user_id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	LoginDate time.Time `json:"login_date"`
}

func NewUser(e *db.User) *User {
	return &User{
		UserID: e.UUID,
		Name:   e.Name,
		Email:  e.Email,
	}
}

type UserWithToken struct {
	User  *User  `json:"user"`
	Token string `json:"token"`
}

func NewUserWithToken(e *db.User, token string) *UserWithToken {
	return &UserWithToken{
		User:  NewUser(e),
		Token: token,
	}
}
