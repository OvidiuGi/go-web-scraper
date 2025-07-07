package parser

import (
	"log"
	"strings"

	"github.com/OvidiuGi/go-web-scraper/internal/shared/model"
	"github.com/gocolly/colly"
)

func ConfigureMainCollector(c *colly.Collector, childCollector *colly.Collector, setting model.ScraperSettings, data *[]model.Data) {
	c.OnRequest(func(r *colly.Request) {
		log.Println("Visiting", r.URL)
	})

	c.OnHTML(setting.SourceSearchTag, func(e *colly.HTMLElement) {
		href := e.Attr("href")

		if setting.VisitChild {
			if strings.Contains(setting.Source, "euronews") {
				if strings.Contains(href, "/articole") {
					childCollector.Visit("https://www.euronews.ro" + href)

					childCollector.OnHTML(setting.ChildSettings.SearchAttr, ChildOnHTMLCallback(setting.ChildSettings, data))
				}
			} else {
				childCollector.Visit(setting.Source + href)

				childCollector.OnHTML(setting.ChildSettings.SearchAttr, ChildOnHTMLCallback2(setting.ChildSettings, data))
			}
		}
	})
}

func ChildOnHTMLCallback(childSettings model.ChildSettings, data *[]model.Data) func(*colly.HTMLElement) {
	return func(e *colly.HTMLElement) {
		var contentParts []string
		e.ForEach("div[itemprop='articleBody'] p", func(i int, el *colly.HTMLElement) {
			text := strings.TrimSpace(el.Text)
			if text != "" {
				contentParts = append(contentParts, text)
			}
		})

		content := strings.Join(contentParts, "\n")

		article := model.Data{
			Title:   e.ChildText(childSettings.TitleAttr),
			URL:     e.Request.URL.String(),
			Content: content,
		}

		*data = append(*data, article)
	}
}

func ChildOnHTMLCallback2(childSettings model.ChildSettings, data *[]model.Data) func(*colly.HTMLElement) {
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

		article := model.Data{
			Title:   e.ChildText(childSettings.TitleAttr),
			URL:     e.Request.URL.String(),
			Content: content,
		}

		*data = append(*data, article)
	}
}
