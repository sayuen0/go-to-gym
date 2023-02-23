package auth

import "github.com/gin-gonic/gin"

// Handlers defines the methods for handling HTTP requests related to user authentication and authorization.
type Handlers interface {
	// Register handles the HTTP request for user registration.
	Register() gin.HandlerFunc
	// Login handles the HTTP request for user login.
	Login() gin.HandlerFunc
	// Logout handles the HTTP request for user logout.
	Logout() gin.HandlerFunc
	// GetUsers handles the HTTP request for retrieving all users.
	GetUsers() gin.HandlerFunc
	// GetUserByID handles the HTTP request for retrieving a user by ID.
	GetUserByID() gin.HandlerFunc
	// GetMe handles the HTTP request for retrieving the authenticated user.
	GetMe() gin.HandlerFunc
	// Update handles the HTTP request for updating the authenticated user.
	Update() gin.HandlerFunc
	// Delete handles the HTTP request for deleting the authenticated user.
	Delete() gin.HandlerFunc
}
