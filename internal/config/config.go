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
}

var Cfg *Config

func LoadConfig() {
	err := godotenv.Load()
	if err != nil {
		log.Println("⚠️ Warning: .env file not found, using system env")
	}

	Cfg = &Config{
		MysqlDSN:  os.Getenv("MYSQL_DSN"),
		RedisAddr: os.Getenv("REDIS_ADDR"),
		RedisPass: os.Getenv("REDIS_PASSWORD"),
		AppPort:   os.Getenv("APP_PORT"),
	}
}
