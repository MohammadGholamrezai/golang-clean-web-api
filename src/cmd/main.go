package main

import (
	"context"
	"log"

	api "github.com/MohammadGholamrezai/golang-clean-web-api/api"
	"github.com/MohammadGholamrezai/golang-clean-web-api/config"
	"github.com/MohammadGholamrezai/golang-clean-web-api/data/cache"
	"github.com/MohammadGholamrezai/golang-clean-web-api/data/db"
)

func main() {
	cfg := config.GetConfig()
	ctx := context.Background()

	err := cache.InitRedis(cfg, ctx)
	if err != nil {
		log.Fatal(err)
	}
	defer cache.CloseRedis()

	err = db.InitDb(cfg)
	if err != nil {
		log.Fatal(err)
	}
	defer db.CloseDb()

	api.InitServer(cfg)
}
