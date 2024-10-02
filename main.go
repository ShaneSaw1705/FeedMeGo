package main

import (
	"feed-me/initializers"
	"feed-me/routes"
	"os"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func init() {
	initializers.Env()
	initializers.ConnectDB()
	initializers.Migrate()
}

func main() {
	r := gin.Default()
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5173"}, // Svelte app URL
		AllowMethods:     []string{"GET", "POST", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true, // Allow cookies to be sent
		MaxAge:           12 * time.Hour,
	}))

	routes.RegisterRoutes(r)

	r.Run(os.Getenv("port"))
}
