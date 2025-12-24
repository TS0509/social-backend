package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"social-backend/internal/config"
	"social-backend/internal/database"
	"social-backend/internal/model"
	"social-backend/internal/router"
)

func main() {
	config.LoadConfig()

	database.InitMySQL()
	database.InitRedis()
	database.DB.AutoMigrate(&model.User{})

	r := router.SetupRouter()

	srv := &http.Server{
		Addr:    ":" + config.Cfg.AppPort,
		Handler: r,
	}

	go func() {
		log.Println("🚀 Server running on port " + config.Cfg.AppPort)
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("❌ Server error: %v", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit

	log.Println("🛑 Shutting down server...")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	srv.Shutdown(ctx)
	log.Println("✅ Server exited")
}
