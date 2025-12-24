package database

import (
	"context"
	"log"
	"time"

	"social-backend/internal/config"

	"github.com/redis/go-redis/v9"
)

var Rdb *redis.Client

func InitRedis() {
	if config.Cfg == nil {
		log.Fatal("❌ Config not loaded")
	}

	Rdb = redis.NewClient(&redis.Options{
		Addr:         config.Cfg.RedisAddr,
		Password:     config.Cfg.RedisPass,
		DB:           0,
		PoolSize:     30,
		MinIdleConns: 5,
	})

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	if err := Rdb.Ping(ctx).Err(); err != nil {
		log.Fatalf("❌ Redis connection failed: %v", err)
	}

	log.Println("✅ Redis connected")
}
