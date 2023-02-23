package main

import (
	"log"
	"os"

	"github.com/sayuen0/go-to-gym/config"
	"github.com/sayuen0/go-to-gym/internal/infrastructure/logger"
	"github.com/sayuen0/go-to-gym/internal/server"
	"github.com/sayuen0/go-to-gym/pkg/db/mysql"
	"github.com/sayuen0/go-to-gym/pkg/db/redis"
	"github.com/sayuen0/go-to-gym/pkg/utils"
)

// @title Go To Gym API
// @version 1.0
// @contact.name sayuen0
// @contact.url https://github.com/sayuen0
// @basePath /
func main() {
	configPath := utils.GetConfigPath(os.Getenv("config"))
	viperConfig, err := config.LoadConfig(configPath)
	if err != nil {
		log.Fatalf("LoadConfig: %v", err)
	}
	cfg, err := config.ParseConfig(viperConfig)
	if err != nil {
		log.Fatalf("ParseConfig: %v", err)
	}
	zl, err := logger.NewLogger()
	if err != nil {
		log.Fatalf("NewLogger: %v", err)
	}

	db := mysql.GetConnection(cfg)
	defer db.Close()

	redisClient := redis.NewRedisClient(cfg)
	defer redisClient.Close()

	s := server.NewServer(cfg, zl, db, redisClient)
	if err := s.Run(); err != nil {
		panic(err)
	}
}
