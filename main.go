package main

import (
	"feed-me/initializers"
	"feed-me/routes"
	"os"

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
		AllowOrigins:     []string{os.Getenv("FrontEndUrl")},
		AllowCredentials: true,
		// other configurations
	}))
	routes.RegisterRoutes(r)

	r.Run(os.Getenv("port"))
}
