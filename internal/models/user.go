package models

import (
	"github.com/sayuen0/go-to-gym/internal/models/db"
	"golang.org/x/crypto/bcrypt"
	"strings"
)

// ---------------------------------------------------------------------------------------------------------------------
// user create request

type UserCreateRequest struct {
	Name           string `json:"name" validate:"required"`
	Email          string `json:"email" validate:"required"`
	Password       string `json:"password" validate:"required"`
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
// user login request

type UserLoginRequest struct {
	Email    string `json:"email" validate:"required"`
	Password string `json:"password" validate:"required"`
}

// ---------------------------------------------------------------------------------------------------------------------
// user update request

type UserUpdateRequest struct {
	UserID         string  `json:"user_id" validate:"required"`
	Name           string  `json:"name" validate:"required"`
	Email          string  `json:"email" validate:"required"`
	Password       string  `json:"password" validate:"required"`
	HashedPassword string  `json:"-"`
	Gender         *string `json:"gender,omitempty"`
}

// TODO: write validate tag

// ---------------------------------------------------------------------------------------------------------------------
// user get response

type User struct {
	UserID string `json:"user_id"`
	Name   string `json:"name"`
	Email  string `json:"email"`
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

// ---------------------------------------------------------------------------------------------------------------------
// users list response

type UsersList struct {
	Paging `json:"paging"`
	Users  []*User `json:"users"`
}

func NewUsersList(users []*db.User, p Paging) *UsersList {
	list := &UsersList{Users: make([]*User, 0, len(users)), Paging: p}
	for _, u := range users {
		list.Users = append(list.Users, NewUser(u))
	}
	return list
}

// ---------------------------------------------------------------------------------------------------------------------
// utils

func CompareUserPassword(inputPassword, hashedPassword string) error {
	if err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(inputPassword)); err != nil {
		return err
	}
	return nil
}
