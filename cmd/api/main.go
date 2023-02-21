package main

import (
	"github.com/sayuen0/go-to-gym/config"
	"github.com/sayuen0/go-to-gym/internal/infrastructure/logger"
	"github.com/sayuen0/go-to-gym/internal/server"
	"github.com/sayuen0/go-to-gym/pkg/utils"
	"log"
	"os"
)

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

	s := server.NewServer(cfg, zl)
	if s.Run(); err != nil {
		log.Fatal(err)
	}
}
