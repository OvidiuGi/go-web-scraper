package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/OvidiuGi/go-web-scraper/models"
	"github.com/OvidiuGi/go-web-scraper/scraper"
	"github.com/gin-gonic/gin"
	"sync"
)

func Scrape(c *gin.Context) {
	var req models.ScrapeRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": "Invalid request format"})
		return
	}

	c.Header("Content-Type", "text/event-stream")
	c.Header("Cache-Control", "no-cache")
	c.Header("Connection", "keep-alive")
	c.Header("Access-Control-Allow-Origin", "*")

	updateChan := make(chan models.ScrapeResponse, len(req.Settings))

	go func() {
		defer close(updateChan)
		scrapeWithUpdates(req.Settings, updateChan)
	}()

	for update := range updateChan {
		data, _ := json.Marshal(update)
		fmt.Fprintf(c.Writer, "data: %s\n\n", data)
		c.Writer.Flush()
	}
}

func scrapeWithUpdates(settings []models.ScraperSettings, updateChan chan<- models.ScrapeResponse) {
	var wg sync.WaitGroup

	for _, setting := range settings {
		wg.Add(1)

		updateChan <- models.ScrapeResponse{
			Source:  setting.Source,
			Status:  "Started",
			Message: "Starting to scrape...",
		}

		go func(setting models.ScraperSettings) {
			defer wg.Done()

			articles := scraper.ScrapeFromSource(setting)

			if len(articles) > 0 {
				updateChan <- models.ScrapeResponse{
					Source:  setting.Source,
					Status:  "Completed",
					Count:   len(articles),
					Data:    articles,
					Message: fmt.Sprintf("Successfully scraped %d articles", len(articles)),
				}
			} else {
				updateChan <- models.ScrapeResponse{
					Source:  setting.Source,
					Status:  "Failed",
					Message: "No articles found or scraping failed",
				}
			}
		}(setting)
	}

	wg.Wait()

	updateChan <- models.ScrapeResponse{
		Status:  "Finished",
		Message: "Scraping completed for all sources",
	}
}
