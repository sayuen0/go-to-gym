package server

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func (s *server) Handle(r *gin.Engine) error {
	// TODO: init repositories

	// TODO init use cases

	// TODO use middlewares

	r.GET("/", func(c *gin.Context) {
		// TODO: identify health check ID
		s.lg.Info("Health check")
		c.String(http.StatusOK, "OK")
	})
	return nil
}
