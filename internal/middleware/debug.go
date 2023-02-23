package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/sayuen0/go-to-gym/internal/infrastructure/logger"
	"net/http"
	"net/http/httputil"
)

// DebugMiddleware returns a middleware function that logs the outgoing request's HTTP dump for debugging purposes.
func (mw *Wrapper) DebugMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		dump, err := httputil.DumpRequestOut(c.Request, true)
		if err != nil {
			c.Status(http.StatusInternalServerError)
			return
		}

		mw.lg.Info("Request dump",
			logger.ByteString("r", dump))
		c.Next()
	}
}
