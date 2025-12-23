package database

import (
	"context"
	"log"
	"social-backend/internal/config"

	"github.com/redis/go-redis/v9"
)

var Rdb *redis.Client

func InitRedis() {
	Rdb = redis.NewClient(&redis.Options{
		Addr:     config.Cfg.RedisAddr,
		Password: config.Cfg.RedisPass,
		DB:       0,
	})

	err := Rdb.Ping(context.Background()).Err()
	if err != nil {
		log.Fatalf("❌ Redis connection failed: %v", err)
	}

	log.Println("✅ Redis connected")
}
