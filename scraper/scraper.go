package scraper

import (
	"github.com/OvidiuGi/go-web-scraper/models"
	"github.com/gocolly/colly"
	"log"
	"strings"
)

func ScrapeFromSource(setting models.ScraperSettings) []models.Data {
	data := []models.Data{}

	c := colly.NewCollector()
	childCollector := c.Clone()

	// TODO: Add allowed domains from requested sources.

	c.OnError(func(_ *colly.Response, err error) {
		log.Println("Something went wrong:", err)
	})

	c.OnHTML(setting.SourceSearchTag, func(e *colly.HTMLElement) {
		href := e.Attr("href")

		if setting.VisitChild {
			if strings.Contains(setting.Source, "euronews") {
				if strings.Contains(href, "/articole") {
					childCollector.Visit("https://www.euronews.ro" + href)

					childCollector.OnHTML(setting.ChildSettings.SearchAttr, ChildOnHTMLCallback(setting.ChildSettings, &data))
				}
			} else {
				childCollector.Visit(setting.Source + href)

				childCollector.OnHTML(setting.ChildSettings.SearchAttr, ChildOnHTMLCallback2(setting.ChildSettings, &data))
			}
		}
	})

	c.Visit(setting.Source)

	return data
}

func ChildOnHTMLCallback(childSettings models.ChildSettings, data *[]models.Data) func(*colly.HTMLElement) {
	return func(e *colly.HTMLElement) {
		var contentParts []string
		// The article content is inside: <div itemprop="articleBody">
		e.ForEach("div[itemprop='articleBody'] p", func(i int, el *colly.HTMLElement) {
			text := strings.TrimSpace(el.Text)
			if text != "" {
				contentParts = append(contentParts, text)
			}
		})

		content := strings.Join(contentParts, "\n")

		article := models.Data{
			Title:   e.ChildText(childSettings.TitleAttr),
			URL:     e.Request.URL.String(),
			Content: content,
		}

		*data = append(*data, article)
	}
}

func ChildOnHTMLCallback2(childSettings models.ChildSettings, data *[]models.Data) func(*colly.HTMLElement) {
	return func(e *colly.HTMLElement) {
		var contentParts []string
		e.ForEach("div.data-app-meta-article > *", func(i int, el *colly.HTMLElement) {
			if el.DOM.Is("p[data-index]") {
				text := strings.TrimSpace(el.Text)
				if text != "" {
					contentParts = append(contentParts, text)
				}
			}

			if el.DOM.Is("ul") {
				var listItems []string
				el.ForEach("li", func(j int, li *colly.HTMLElement) {
					text := strings.TrimSpace(li.Text)
					if text != "" {
						listItems = append(listItems, "• "+text)
					}
				})
				if len(listItems) > 0 {
					contentParts = append(contentParts, strings.Join(listItems, "\n"))
				}
			}
		})

		content := strings.Join(contentParts, "\n")

		article := models.Data{
			Title:   e.ChildText(childSettings.TitleAttr),
			URL:     e.Request.URL.String(),
			Content: content,
		}

		*data = append(*data, article)
	}
}
