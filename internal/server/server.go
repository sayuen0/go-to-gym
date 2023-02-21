package server

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/sayuen0/go-to-gym/config"
	"github.com/sayuen0/go-to-gym/internal/infrastructure/logger"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

const (
	_certFile  = "ssl/server.crt"
	_keyFile   = "ssl/server.key"
	ctxTimeout = 5
)

type Server interface {
	Run() error
}

type server struct {
	gin    *gin.Engine
	cfg    *config.Config
	logger logger.Logger
}

func NewServer(
	cfg *config.Config, lg logger.Logger,
) Server {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Hello world")
	})
	return &server{
		gin:    r,
		cfg:    cfg,
		logger: lg,
	}
}

func (s *server) Run() error {
	srv := &http.Server{
		Addr:    s.cfg.Server.Port,
		Handler: s.gin,
	}
	if s.cfg.Server.SSL {
		// TODO: map handlers
		// TODO: set read & write timeout

		go func() {
			s.logger.Info("Server is listening", logger.String("port", s.cfg.Server.Port))
			if err := srv.ListenAndServeTLS(_certFile, _keyFile); err != nil {
				s.logger.Fatal("Error starting TLS server", logger.Error(err))
			}
		}()
	} else {
		// TODO: map handlers
		// TODO: set read & write timeout

		go func() {
			s.logger.Info("Server is listening", logger.String("port", s.cfg.Server.Port))
			if err := srv.ListenAndServe(); err != nil {
				s.logger.Fatal("Error starting server", logger.Error(err))
			}
		}()
	}

	// graceful shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)

	<-quit
	ctx, shutdown := context.WithTimeout(context.Background(), ctxTimeout*time.Second)
	defer shutdown()
	s.logger.Info("Server is Shutting down")
	return srv.Shutdown(ctx)
}
