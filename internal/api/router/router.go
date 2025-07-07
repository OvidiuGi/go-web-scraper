package router

import (
	"github.com/OvidiuGi/go-web-scraper/internal/api/handler"

	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	r := gin.Default()

	r.StaticFile("/", "./static/index.html")

	apiV1 := r.Group("/api/v1")
	{
		apiV1.POST("/scrape", handler.Scrape)
	}

	return r
}
