package server

import (
	"context"
	"database/sql"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
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
	gin         *gin.Engine
	cfg         *config.Config
	lg          logger.Logger
	db          *sql.DB
	redisClient *redis.Client
}

func NewServer(
	cfg *config.Config, lg logger.Logger, db *sql.DB, redisClient *redis.Client,
) Server {
	r := gin.Default()
	return &server{
		gin:         r,
		cfg:         cfg,
		lg:          lg,
		db:          db,
		redisClient: redisClient,
	}
}

func (s *server) Run() error {
	if err := s.Handle(s.gin); err != nil {
		s.lg.Error("Handle error", logger.Error(err))
	}
	srv := &http.Server{
		Addr:    s.cfg.Server.Port,
		Handler: s.gin,
	}
	if s.cfg.Server.SSL {
		// TODO: set read & write timeout

		go func() {
			s.lg.Info("TLS Server is listening", logger.String("port", s.cfg.Server.Port))
			if err := srv.ListenAndServeTLS(_certFile, _keyFile); err != nil {
				s.lg.Fatal("Error starting TLS server", logger.Error(err))
			}
		}()
	} else {
		// TODO: set read & write timeout

		go func() {
			s.lg.Info("Server is listening", logger.String("port", s.cfg.Server.Port))
			if err := srv.ListenAndServe(); err != nil {
				s.lg.Fatal("Error starting server", logger.Error(err))
			}
		}()
	}

	// graceful shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)

	<-quit
	ctx, shutdown := context.WithTimeout(context.Background(), ctxTimeout*time.Second)
	defer shutdown()
	s.lg.Info("Server is Shutting down")
	return srv.Shutdown(ctx)
}
