package main

import (
	"log"
	"net/http"

	"go-user-auth/internal/config"
	"go-user-auth/internal/database"

	"github.com/gin-gonic/gin"
)

func main() {
	config := config.LoadConfig()
	_, err := database.Connect(config)
	if err != nil {
		log.Fatal("Failed to connect to the database:", err)
	}

	router := gin.Default()
	router.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "ok"})
	})

	log.Printf("🚀 Server running on http://localhost:%s", config.AppPort)

	if err := router.Run(":" + config.AppPort); err != nil {
		log.Fatal("Failed to start server: ", err)
	}
}
