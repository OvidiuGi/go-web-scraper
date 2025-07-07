package main

import (
	"log"
	"net/http"

	"github.com/OvidiuGi/go-web-scraper/internal/api/router"
	"github.com/OvidiuGi/go-web-scraper/internal/shared/config"
)

func main() {
	cfg, err := config.LoadConfig()

	if err != nil {
		panic("Failed to load configuration: " + err.Error())
	}

	r := router.NewRouter()

	log.Printf("API server starting on port %s", cfg.ApiPort)
	if err := http.ListenAndServe(":"+cfg.ApiPort, r); err != nil {
		log.Fatalf("Failed to start API server: %v", err)
	}
}
