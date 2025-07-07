package core

import (
	"log"
	"net/url"

	"github.com/OvidiuGi/go-web-scraper/internal/scraper/parser"
	"github.com/OvidiuGi/go-web-scraper/internal/shared/model"
	"github.com/gocolly/colly"
)

func ScrapeFromSource(setting model.ScraperSettings) []model.Data {
	data := []model.Data{}

	c := colly.NewCollector()
	childCollector := c.Clone()

	domain, err := url.Parse(setting.Source)

	if err != nil {
		log.Println("Invalid source URL:", err)
		return data
	}

	c.AllowedDomains = []string{
		domain.Hostname(),
	}

	c.OnError(func(_ *colly.Response, err error) {
		log.Println("Something went wrong:", err)
	})

	parser.ConfigureMainCollector(c, childCollector, setting, &data)

	c.Visit(setting.Source)

	return data
}
