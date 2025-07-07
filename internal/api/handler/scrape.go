package handler

import (
	"encoding/json"
	"fmt"
	"sync"

	"github.com/OvidiuGi/go-web-scraper/internal/scraper/core"
	"github.com/OvidiuGi/go-web-scraper/internal/shared/model"
	"github.com/gin-gonic/gin"
)

func Scrape(c *gin.Context) {
	var req model.ScrapeRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": "Invalid request format"})
		return
	}

	c.Header("Content-Type", "text/event-stream")
	c.Header("Cache-Control", "no-cache")
	c.Header("Connection", "keep-alive")
	c.Header("Access-Control-Allow-Origin", "*")

	updateChan := make(chan model.ScrapeResponse, len(req.Settings))

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

func scrapeWithUpdates(settings []model.ScraperSettings, updateChan chan<- model.ScrapeResponse) {
	var wg sync.WaitGroup

	for _, setting := range settings {
		wg.Add(1)

		updateChan <- model.ScrapeResponse{
			Source:  setting.Source,
			Status:  "Started",
			Message: "Starting to scrape...",
		}

		go func(setting model.ScraperSettings) {
			defer wg.Done()

			data := core.ScrapeFromSource(setting)

			if len(data) > 0 {
				updateChan <- model.ScrapeResponse{
					Source:  setting.Source,
					Status:  "Completed",
					Count:   len(data),
					Data:    data,
					Message: fmt.Sprintf("Successfully scraped %d articles", len(data)),
				}
			} else {
				updateChan <- model.ScrapeResponse{
					Source:  setting.Source,
					Status:  "Failed",
					Message: "No articles found or scraping failed",
				}
			}
		}(setting)
	}

	wg.Wait()

	updateChan <- model.ScrapeResponse{
		Status:  "Finished",
		Message: "Scraping completed for all sources",
	}
}
