package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/thewolmer/proxy-api-server/utils"
)

func TMDbProxy(c *gin.Context) {
	proxy := utils.CreateReverseProxy("https://api.themoviedb.org/3", "/v1/tmdb")
	proxy.ServeHTTP(c.Writer, c.Request)
}

func GitHubProxy(c *gin.Context) {
	proxy := utils.CreateReverseProxy("https://api.github.com", "/v1/github")
	proxy.ServeHTTP(c.Writer, c.Request)
}
