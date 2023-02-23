package models

import (
	"github.com/sayuen0/go-to-gym/internal/models/db"
	"strings"
)

// ---------------------------------------------------------------------------------------------------------------------
// user create request

// TODO: パスワード周り全部外出しした方がいい

// UserCreateRequest represents user create/register request
type UserCreateRequest struct {
	Name     string `json:"name" validate:"required"`
	Email    string `json:"email" validate:"required"`
	Password string `json:"password" validate:"required"`
}

// PrepareForCreate formats its auth info
func (u *UserCreateRequest) PrepareForCreate() error {
	u.Email = strings.ToLower(strings.TrimSpace(u.Email))
	u.Password = strings.TrimSpace(u.Password)
	return nil
}

// ---------------------------------------------------------------------------------------------------------------------
// user login request

// UserLoginRequest represents user login request
type UserLoginRequest struct {
	Email    string `json:"email" validate:"required"`
	Password string `json:"password" validate:"required"`
}

// ---------------------------------------------------------------------------------------------------------------------
// user update request

// UserUpdateRequest represents user update request
type UserUpdateRequest struct {
	UserID         string `json:"user_id" validate:"required"`
	Name           string `json:"name" validate:"required"`
	Email          string `json:"email" validate:"required"`
	Password       string `json:"password" validate:"required"`
	HashedPassword string `json:"-"`
}

// TODO: write validate tag

// ---------------------------------------------------------------------------------------------------------------------
// user get response

// User represents a response model of user
type User struct {
	UserID string `json:"user_id"`
	Name   string `json:"name"`
	Email  string `json:"email"`
}

// NewUser is a constructor of User
func NewUser(e *db.User) *User {
	return &User{
		UserID: e.UserID,
		Name:   e.Name,
		Email:  e.Email,
	}
}

// UserWithToken is a wrapper of User with Token
type UserWithToken struct {
	User  *User  `json:"user"`
	Token string `json:"token"`
}

// NewUserWithToken is a constructor of UserWithToken
func NewUserWithToken(e *db.User, token string) *UserWithToken {
	return &UserWithToken{
		User:  NewUser(e),
		Token: token,
	}
}

// ---------------------------------------------------------------------------------------------------------------------
// users list response

// UsersList represents a response model of users list with paging information
type UsersList struct {
	Paging `json:"paging"`
	Users  []*User `json:"users"`
}

// NewUsersList is a constructor of UsersList
func NewUsersList(users []*db.User, p Paging) *UsersList {
	list := &UsersList{Users: make([]*User, 0, len(users)), Paging: p}
	for _, u := range users {
		list.Users = append(list.Users, NewUser(u))
	}
	return list
}
