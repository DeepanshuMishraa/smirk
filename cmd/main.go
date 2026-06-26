package main

import (
	"log"

	"github.com/DeepanshuMishraa/vid-processing-go.git/config"
	"github.com/DeepanshuMishraa/vid-processing-go.git/db"
	"github.com/gin-gonic/gin"
)

func main() {
	cfg, err := config.Load()

	if err != nil {
		log.Fatalf("Failed to Load Env Variables: %v", err)
	}

	log.Println("Loaded Env Variables")

	_, err = db.Connect(cfg.DATABASE_URL)

	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	router := gin.Default()
	router.SetTrustedProxies(nil)

	router.GET("/health", gin.HandlerFunc(func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status": "ok",
		})
	}))

	router.Run(":" + cfg.PORT)
}
