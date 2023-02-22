package utils

import (
	"github.com/gin-gonic/gin"
	"github.com/sayuen0/go-to-gym/config"
	"github.com/sayuen0/go-to-gym/internal/infrastructure/logger"
)

func GetIpAddress(c *gin.Context) string {
	return c.RemoteIP()
}

func LogResponseError(c *gin.Context, lg logger.Logger, err error) {
	lg.Error("Error response with log",
		// TODO: get request id
		logger.String("ip_address", GetIpAddress(c)),
		logger.Error(err),
	)
}

func CreateSessionCookie(c *gin.Context, cfg *config.Config, value string) {
	c.SetCookie(
		cfg.Session.Name,
		value,
		cfg.Session.Expire,
		"/",
		"",
		cfg.Cookie.Secure,
		cfg.Cookie.HTTPOnly,
	)
}

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

func ReadRequest(c *gin.Context, request any) error {
	if err := c.Bind(request); err != nil {
		return err
	}
	return ValidateStruct(c.Request.Context(), request)
}

func GetConfigPath(configPath string) string {
	return "./config/config-local"
}
