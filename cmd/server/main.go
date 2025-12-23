package main

import (
	"log"
	"social-backend/internal/config"
	"social-backend/internal/database"
	"social-backend/internal/model"
	"social-backend/internal/router"
)

func main() {
	config.LoadConfig()

	// Initialize Database & Redis
	database.InitMySQL()
	database.InitRedis()

	// AutoMigrate — create table if not exist
	err := database.DB.AutoMigrate(&model.User{})
	if err != nil {
		log.Fatalf("❌ Migration failed: %v", err)
	}

	r := router.SetupRouter()

	log.Println("🚀 Server running at http://localhost:" + config.Cfg.AppPort)
	r.Run(":" + config.Cfg.AppPort)
}
