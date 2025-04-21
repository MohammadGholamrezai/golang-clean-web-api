package main

import (
	"context"

	api "github.com/MohammadGholamrezai/golang-clean-web-api/api"
	"github.com/MohammadGholamrezai/golang-clean-web-api/config"
	"github.com/MohammadGholamrezai/golang-clean-web-api/data/cache"
	"github.com/MohammadGholamrezai/golang-clean-web-api/data/db"
	"github.com/MohammadGholamrezai/golang-clean-web-api/pkg/logging"
)

func main() {
	cfg := config.GetConfig()
	ctx := context.Background()

	logger := logging.NewLogger(cfg)

	err := cache.InitRedis(cfg, ctx)
	if err != nil {
		logger.Fatal(logging.Redis, logging.Startup, err.Error(), nil)
	}
	defer cache.CloseRedis()

	err = db.InitDb(cfg)
	if err != nil {
		logger.Fatal(logging.Postgres, logging.Startup, err.Error(), nil)
	}
	defer db.CloseDb()

	api.InitServer(cfg)
}
