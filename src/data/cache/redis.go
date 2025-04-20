package cache

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/MohammadGholamrezai/golang-clean-web-api/config"
	"github.com/redis/go-redis/v9"
)

var redisClient *redis.Client

func InitRedis(cfg *config.Config, ctx context.Context) error {
	redisClient = redis.NewClient(&redis.Options{
		Addr:            fmt.Sprintf("%s:%s", cfg.Redis.Host, cfg.Redis.Port),
		Password:        cfg.Redis.Password,
		DB:              cfg.Redis.Db,
		DialTimeout:     cfg.Redis.DialTimeout * time.Second,
		ReadTimeout:     cfg.Redis.ReadTimeout * time.Second,
		WriteTimeout:    cfg.Redis.WriteTimeout * time.Second,
		PoolSize:        cfg.Redis.PoolSize,
		PoolTimeout:     cfg.Redis.PoolTimeout * time.Second,
		MinIdleConns:    10,
		ConnMaxIdleTime: 5 * time.Minute,
		ConnMaxLifetime: 30 * time.Minute,
	})

	val, err := redisClient.Ping(ctx).Result()
	if err != nil {
		log.Fatalf("Ping error: %v", err)
	}

	fmt.Println("Ping response:", val)
	return nil
}

func GetRedis() *redis.Client {
	return redisClient
}

func CloseRedis() {
	redisClient.Close()
}
