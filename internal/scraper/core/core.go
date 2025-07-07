package core

import (
	"log"

	"github.com/OvidiuGi/go-web-scraper/internal/scraper/parser"
	"github.com/OvidiuGi/go-web-scraper/internal/shared/model"
	"github.com/gocolly/colly"
)

func ScrapeFromSource(setting model.ScraperSettings) []model.Data {
	data := []model.Data{}

	c := colly.NewCollector()
	childCollector := c.Clone()

	c.OnError(func(_ *colly.Response, err error) {
		log.Println("Something went wrong:", err)
	})

	parser.ConfigureMainCollector(c, childCollector, setting, &data)

	c.Visit(setting.Source)

	return data
}
