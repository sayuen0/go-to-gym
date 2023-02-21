package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
)

func main() {
	router := gin.Default()
	router.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Hello world")
	})
	crt := os.Getenv("SERVER_CRT")
	key := os.Getenv("SERVER_KEY")
	router.RunTLS(":8080", crt, key)
}
