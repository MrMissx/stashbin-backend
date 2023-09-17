package main

import (
	"log"

	"github.com/gin-contrib/cors"
	"github.com/mrmissx/stashbin-backend/controllers"
	"github.com/mrmissx/stashbin-backend/models"
	"github.com/mrmissx/stashbin-backend/utils"

	"github.com/gin-gonic/gin"
	_ "github.com/joho/godotenv/autoload"
)

func main() {
	r := gin.New()
	r.SetTrustedProxies([]string{"127.0.0.1", "::1"})
	r.TrustedPlatform = "X-Real-IP"
	r.Use(cors.New(cors.Config{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{"GET", "POST"},
		AllowHeaders: []string{"Origin", "Content-Type"},
	}))
	r.Use(gin.LoggerWithConfig(gin.LoggerConfig{
		SkipPaths: []string{"/api/health"},
	}))
	r.Use(gin.Recovery())

	log.Println("Connecting to database")
	models.Connect()
	log.Println("Connected to database")

	log.Println("Registering routes")
	r.GET("/api/health", controllers.HealthCheck)
	// Document Routes
	r.GET("/api/document", controllers.GetDocumentBySlug)
	r.POST("/api/document", controllers.CreateDocument)

	port := utils.GetEnv("PORT", "8080")
	log.Printf("Listening on port %s", port)
	r.Run()
}
