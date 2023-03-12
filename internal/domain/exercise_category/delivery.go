package exercise_category

import "github.com/gin-gonic/gin"

type Handlers interface {
	Create() gin.HandlerFunc
	//List() gin.HandlerFunc
	//Get() gin.HandlerFunc
	//Update() gin.HandlerFunc
	//Delete() gin.HandlerFunc
}
