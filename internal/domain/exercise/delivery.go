package exercise

import "github.com/gin-gonic/gin"

// Handlers provides HTTP request handlers for creating, listing, getting,
// updating, and deleting items.
type Handlers interface {
	// Create handles a POST request for creating a new item.
	Create() gin.HandlerFunc

	//// List handles a GET request for listing all items.
	//List() gin.HandlerFunc
	//
	//// Get handles a GET request for getting a specific item.
	//Get() gin.HandlerFunc
	//
	//// Update handles a PUT request for updating an existing item.
	//Update() gin.HandlerFunc
	//
	//// Delete handles a DELETE request for deleting an item.
	//Delete() gin.HandlerFunc
}
