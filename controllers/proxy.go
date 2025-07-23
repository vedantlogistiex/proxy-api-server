package controllers

import (
	"net/http"
	// "net/url"

	"github.com/gin-gonic/gin"
	"github.com/thewolmer/proxy-api-server/utils"
	"os"
)

func TMDbProxy(c *gin.Context) {
	apiKey := os.Getenv("TMDB_API_KEY")
	if apiKey == "" {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "TMDB_API_KEY not set"})
		return
	}

	println("👉 Incoming request path:", c.Request.URL.Path)

	query := c.Request.URL.Query()
	query.Set("api_key", apiKey)
	c.Request.URL.RawQuery = query.Encode()

	proxy := utils.CreateReverseProxy("https://api.themoviedb.org/3", "/v1/tmdb")
	proxy.ServeHTTP(c.Writer, c.Request)
}


func GitHubProxy(c *gin.Context) {
	proxy := utils.CreateReverseProxy("https://api.github.com", "/v1/github")
	proxy.ServeHTTP(c.Writer, c.Request)
}
