package main

import (
	"github.com/OvidiuGi/go-web-scraper/handlers"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.StaticFile("/", "./static/index.html")

	router.POST("/scrape", handlers.Scrape)

	router.Run(":8080")
}
