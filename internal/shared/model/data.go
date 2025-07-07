package model

import "time"

// Data represents a scraped article with its metadata
type Data struct {
	Source      string
	Title       string
	URL         string
	Content     string
	Summary     string
	PublishedAt time.Time
}
