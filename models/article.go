package models

import "time"

type Data struct {
	Source      string
	Title       string
	URL         string
	Content     string
	Summary     string
	PublishedAt time.Time
}

type ScrapeRequest struct {
	Settings []ScraperSettings `json:"settings,omitempty"`
}

type ScraperSettings struct {
	Source string `json:"source,omitempty"`
	// Concurrency     int    `json:"concurrency,omitempty"`
	SourceSearchTag string        `json:"source_search_tag,omitempty"`
	VisitChild      bool          `json:"visit_child,omitempty"`
	ChildSettings   ChildSettings `json:"child_settings,omitempty"`
}

type ChildSettings struct {
	SearchAttr string `json:"search_attr,omitempty"`
	TitleAttr  string `json:"title_attr,omitempty"`
}

type ScrapeResponse struct {
	Source  string `json:"source"`
	Status  string `json:"status"`
	Count   int    `json:"count"`
	Message string `json:"message,omitempty"`
	Data    []Data `json:"data,omitempty"`
}
