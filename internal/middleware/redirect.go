package middleware

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func (mw *Wrapper) RedirectToHTTPS() gin.HandlerFunc {
	return func(c *gin.Context) {
		if c.Request.TLS == nil {
			// Redirect to HTTPS
			url := "https://" + c.Request.Host + c.Request.URL.Path
			if len(c.Request.URL.RawQuery) > 0 {
				url += "?" + c.Request.URL.RawQuery
			}
			c.Redirect(http.StatusMovedPermanently, url)
			c.Abort()
		}
	}
}
