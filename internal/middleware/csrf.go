package middleware

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sayuen0/go-to-gym/internal/infrastructure/logger"
	"github.com/sayuen0/go-to-gym/pkg/csrf"
	"github.com/sayuen0/go-to-gym/pkg/httperrors"
)

// CSRF requires request has csrf token
func (mw *Wrapper) CSRF() gin.HandlerFunc {
	return func(c *gin.Context) {
		if !mw.cfg.Server.CSRF {
			c.Next()

			return
		}

		token := c.Request.Header.Get(csrf.CSRFHeader)
		if token == "" {
			mw.lg.Error("CSRF Middleware get CSRF header",
				logger.String("token", token),
				logger.Error(errors.New("empty CSRF token")))
			c.JSON(http.StatusForbidden, httperrors.NewRestError(http.StatusForbidden, "Invalid CSRF token", "no CSRF token"))

			return
		}

		sid, found := c.Get("sid")
		if !found {
			c.JSON(http.StatusUnauthorized, httperrors.NewRestError(http.StatusUnauthorized, "sid not found", ""))

			return
		}

		if sessionID, ok := sid.(string); !ok || csrf.ValidToken(token, sessionID, mw.lg) {
			mw.lg.Error("CSRF Middleware csrf.ValidToken",
				logger.String("token", token))
			c.JSON(http.StatusForbidden, httperrors.NewRestError(http.StatusForbidden, "Invalid CSRF token", "no CSRF token"))

			return
		}

		c.Next()
	}
}
