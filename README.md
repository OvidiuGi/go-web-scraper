# Web Scraper

A powerful web scraping application built with Go that can extract articles from news websites with real-time progress tracking via Server-Sent Events (SSE).

## Features

- 🕷️ **Multi-site scraping** - Configurable scrapers for different news websites
- 📡 **Real-time updates** - Stream scraping progress using Server-Sent Events
- 🎯 **Flexible selectors** - CSS selector-based content extraction
- 🚀 **Concurrent processing** - Parallel scraping for multiple sources
- 🌐 **Clean web interface** - Simple UI for managing scraping tasks
- ⚙️ **Configurable settings** - Customizable scraping parameters per source

## Project Structure

```
go-web-scraper/
├── handlers/           # HTTP request handlers
│   ├── health.go      # Health check endpoint
│   └── scrape.go      # Scraping endpoint with SSE
├── models/            # Data models and structures
│   └── article.go     # Article and scraper configuration models
├── scraper/           # Core scraping logic
│   └── scraper.go     # Colly-based web scraping implementation
├── static/            # Frontend assets
│   └── index.html     # Web interface
├── main.go           # Application entry point
├── go.mod            # Go module dependencies
└── README.md         # This file
```

## Dependencies

- **[Gin](https://github.com/gin-gonic/gin)** - HTTP web framework
- **[Colly](https://github.com/gocolly/colly)** - Web scraping framework
- **Frontend** - Vanilla HTML/CSS/JavaScript with streaming support

## Installation

1. **Clone the repository:**
   ```bash
   git clone https://github.com/yourusername/go-web-scraper.git
   cd go-web-scraper
   ```

2. **Install dependencies:**
   ```bash
   go mod download
   ```

3. **Run the application:**
   ```bash
   go run main.go
   ```

4. **Open your browser:**
   ```
   http://localhost:8080
   ```

## API Reference

### POST /scrape

Initiates a scraping session with real-time progress updates via Server-Sent Events.

#### Request Body

```json
{
    "settings": [
        {
            "source": "https://www.euronews.ro/categorii/romania",
            "source_search_tag": "article a",
            "visit_child": true,
            "child_settings": {
                "search_attr": "body",
                "title_attr": "h1"
            }
        },
        {
            "source": "https://www.digi24.ro/stiri/actualitate/politica",
            "source_search_tag": "article h2.article-title a",
            "visit_child": true,
            "child_settings": {
                "search_attr": "body",
                "title_attr": "h1"
            }
        }
    ]
}
```

#### Request Parameters

- **`settings`** (array) - Array of scraper configurations
  - **`source`** (string) - Base URL to scrape
  - **`source_search_tag`** (string) - CSS selector for finding article links
  - **`visit_child`** (boolean) - Whether to visit individual article pages
  - **`child_settings`** (object) - Configuration for article page extraction
    - **`search_attr`** (string) - CSS selector for content container
    - **`title_attr`** (string) - CSS selector for article title

#### Response Format (Server-Sent Events)

The endpoint streams responses in SSE format:

**Started Event:**
```
data: {"source":"https://www.euronews.ro/categorii/romania","status":"Started","message":"Starting to scrape..."}
```

**Completed Event:**
```
data: {"source":"https://www.euronews.ro/categorii/romania","status":"Completed","count":30,"message":"Successfully scraped 30 articles","data":[...]}
```

**Finished Event:**
```
data: {"status":"Finished","message":"Scraping completed for all sources"}
```

#### Example using cURL

```bash
curl -X POST http://localhost:8080/scrape \
  -H "Content-Type: application/json" \
  -d '{
    "settings": [
        {
            "source": "https://www.euronews.ro/categorii/romania",
            "source_search_tag": "article a",
            "visit_child": true,
            "child_settings": {
                "search_attr": "body",
                "title_attr": "h1"
            }
        }
    ]
}'
```

### GET /health

Health check endpoint.

#### Response

```json
{
    "message": "Hello, World!"
}
```

## Configuration

### Scraper Settings Explained

- **Source Search Tag**: CSS selector used to find article links on listing pages
  - `"article a"` - Any `<a>` tag inside `<article>` elements
  - `"article h2.article-title a"` - Specific title links in articles

- **Child Settings**: Configuration for extracting content from individual articles
  - `search_attr: "body"` - Look for content within the entire page body
  - `title_attr: "h1"` - Extract title from the main heading

### Supported Sites

The scraper is currently configured for:

- **Euronews Romania** (`https://www.euronews.ro/categorii/romania`)
- **Digi24 Politics** (`https://www.digi24.ro/stiri/actualitate/politica`)

Additional sites can be added by configuring appropriate CSS selectors.

## Web Interface

The application includes a user-friendly web interface accessible at `http://localhost:8080`:

- **URL Input**: Add multiple URLs (one per line)
- **Real-time Progress**: See scraping progress as it happens
- **Article Display**: View scraped articles with titles and content previews
- **Error Handling**: Clear error messages for failed scraping attempts

## Development

### Adding New Sites

1. Analyze the target site's HTML structure
2. Identify CSS selectors for:
   - Article links on listing pages
   - Article title on individual pages
   - Article content containers
3. Test selectors using browser developer tools
4. Configure the scraper settings accordingly

### Code Structure

- **Handlers**: HTTP request/response logic with SSE streaming
- **Models**: Data structures for articles and configuration
- **Scraper**: Core scraping logic using Colly framework
- **Frontend**: Vanilla JavaScript with streaming response handling

## License

This project is open source and available under the [MIT License](LICENSE).

## Contributing

1. Fork the repository
2. Create a feature branch
3. Make your changes
4. Add tests if applicable
5. Submit a pull request

## Disclaimer

This tool is for educational purposes. Please respect robots.txt files and website terms of service when scraping. Implement appropriate rate limiting and be respectful to target websites.
