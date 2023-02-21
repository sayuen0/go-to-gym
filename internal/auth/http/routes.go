package http

import (
	"github.com/gin-gonic/gin"
	"github.com/sayuen0/go-to-gym/internal/auth"
)

func HandleUserRoutes(rg *gin.RouterGroup, h auth.Handlers) {
	rg.POST("/register", h.Register())
	rg.POST("/login", h.Login())
	rg.POST("/logout", h.Logout())
	rg.GET("/", h.GetUsers())
	rg.GET("/:user_id", h.GetUserByID())
	rg.Use(mw.AuthSessionMiddleware)
	rg.GET("/me", h.GetMe())
	rg.PUT("/:user_id", h.Update())
	rg.DELETE("/:user_id", h.Delete())
}
