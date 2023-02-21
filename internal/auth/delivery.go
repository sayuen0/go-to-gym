package auth

import "github.com/gin-gonic/gin"

type Handlers interface {
	Register() gin.HandlerFunc
	Login() gin.HandlerFunc
	Logout() gin.HandlerFunc
	GetUsers() gin.HandlerFunc
	GetUserByID() gin.HandlerFunc
	GetMe() gin.HandlerFunc
	Update() gin.HandlerFunc
	Delete() gin.HandlerFunc
}
