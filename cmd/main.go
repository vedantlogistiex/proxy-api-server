package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/thewolmer/proxy-api-server/config"
	"github.com/thewolmer/proxy-api-server/controllers"
	"github.com/thewolmer/proxy-api-server/middlewares"
)

func init() {
	config.LoadEnv()
}

func main() {
	r := gin.Default()
	r.Use(middlewares.CheckHeaderMiddleware())
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Welcome to the Proxy Server",
		})
	})

	// TMDb proxy route
	r.Any("/v1/tmdb/*proxyPath", controllers.TMDbProxy)

	// GitHub proxy route
	r.Any("/v1/github/*proxyPath", controllers.GitHubProxy)

	if err := r.Run(":" + os.Getenv("PORT")); err != nil {
		log.Fatalf("Could not run server: %v", err)
	}
}
