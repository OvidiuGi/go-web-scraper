package model

// ScrapeRequest represents the request payload for the scrape endpoint
type ScrapeRequest struct {
	Settings []ScraperSettings `json:"settings,omitempty"`
}

// ScrapeResponse represents the streaming response from the scrape endpoint
type ScrapeResponse struct {
	Source  string `json:"source"`
	Status  string `json:"status"`
	Count   int    `json:"count"`
	Message string `json:"message,omitempty"`
	Data    []Data `json:"data,omitempty"`
}

// ScraperSettings contains configuration for scraping a specific source
type ScraperSettings struct {
	Source          string        `json:"source,omitempty"`
	SourceSearchTag string        `json:"source_search_tag,omitempty"`
	VisitChild      bool          `json:"visit_child,omitempty"`
	ChildSettings   ChildSettings `json:"child_settings,omitempty"`
}

// ChildSettings contains configuration for extracting data from child pages
type ChildSettings struct {
	SearchAttr string `json:"search_attr,omitempty"`
	TitleAttr  string `json:"title_attr,omitempty"`
}
