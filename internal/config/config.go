package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	MysqlDSN  string
	RedisAddr string
	RedisPass string
	AppPort   string
	JWTSecret string
}

var Cfg *Config

func LoadConfig() {
	// Load .env
	err := godotenv.Load(".env")
	if err != nil {
		log.Println("⚠️ .env not found, using system environment")
	}

	port := os.Getenv("APP_PORT")
	if port == "" {
		port = "8080"
		log.Println("⚠️ APP_PORT missing, using default 8080")
	}

	Cfg = &Config{
		MysqlDSN:  os.Getenv("MYSQL_DSN"),
		RedisAddr: os.Getenv("REDIS_ADDR"),
		RedisPass: os.Getenv("REDIS_PASSWORD"),
		AppPort:   port,
		JWTSecret: os.Getenv("JWT_SECRET"),
	}

	if Cfg.JWTSecret == "" {
		log.Println("⚠️ JWT_SECRET is missing — please set it in .env")
	}
}
