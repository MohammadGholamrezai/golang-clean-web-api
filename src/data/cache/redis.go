package cache

import (
	"context"
	"encoding/json"
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

func Set[T any](c *redis.Client, key string, value T, duration time.Duration) error {
	v, err := json.Marshal(value)
	if err != nil {
		return err
	}

	// if we don't have any context using context.Background()
	return c.Set(context.Background(), key, v, duration).Err()
}

func Get[T any](c *redis.Client, key string) (T, error) {
	var destination T = *new(T)
	v, err := c.Get(context.Background(), key).Result()
	if err != nil {
		return destination, err
	}

	err = json.Unmarshal([]byte(v), &destination)
	if err != nil {
		return destination, err
	}
	return destination, nil
}
