package main

import (
	"feed-me/initializers"
	"feed-me/routes"
	"os"

	"github.com/gin-gonic/gin"
)

func init() {
	initializers.Env()
}

func main() {
	r := gin.Default()

	routes.RegisterRoutes(r)

	r.Run(os.Getenv("port"))
}
