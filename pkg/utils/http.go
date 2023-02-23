package utils

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/sayuen0/go-to-gym/config"
	"github.com/sayuen0/go-to-gym/internal/infrastructure/logger"
	"github.com/sayuen0/go-to-gym/pkg/httperrors"
)

// GetIPAddress returns requesters' remote IP address
func GetIPAddress(c *gin.Context) string {
	return c.RemoteIP()
}

func ErrorResponseWithLog(c *gin.Context, lg logger.Logger, err error) {
	lg.Error("ErrorResponseWithLog",
		// TODO: get request id
		logger.String("IP", GetIPAddress(c)),
		logger.Error(err),
	)
	c.JSON(httperrors.ErrorResponse(err))
}

// LogResponseError writes an error log to logger.Logger
func LogResponseError(c *gin.Context, lg logger.Logger, err error) {
	lg.Error("Error response with log",
		// TODO: get request id
		logger.String("IP", GetIPAddress(c)),
		logger.Error(err),
	)
}

// CreateSessionCookie creates a session to cookie
func CreateSessionCookie(c *gin.Context, cfg *config.Config, value string) {
	c.SetCookie(
		cfg.Session.Name,
		value,
		cfg.Session.Expire,
		cfg.Cookie.Path,
		cfg.Cookie.Domain,
		cfg.Cookie.Secure,
		cfg.Cookie.HTTPOnly,
	)
}

// DeleteSessionCookie deletes the session from cookie
func DeleteSessionCookie(c *gin.Context, cfg *config.Config) {
	c.SetCookie(
		cfg.Session.Name,
		"",
		-1,
		"/",
		"",
		false,
		true)
}

// UserCtxKey is a key user for the User object in the context
type UserCtxKey struct{}

// ReadRequest reads http request body and validates it
func ReadRequest(c *gin.Context, request any) error {
	if err := c.Bind(request); err != nil {
		return err
	}
	return ValidateStruct(c.Request.Context(), request)
}

// GetConfigPath returns application configuration file path
func GetConfigPath(configPath string) string {
	// TODO: set by env
	return "./config/config-local"
}

// GetRequestCtx returns request context.Context
func GetRequestCtx(c *gin.Context) context.Context {
	return c.Request.Context()
}
