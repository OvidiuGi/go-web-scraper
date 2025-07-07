# Go Web Scraper

A powerful web scraping application built with Go that can extract articles from news websites with real-time progress tracking via Server-Sent Events (SSE).

## Features

- 🕷️ **Multi-site scraping** - Configurable scrapers for different news websites
- 📡 **Real-time updates** - Stream scraping progress using Server-Sent Events
- 🎯 **Flexible selectors** - CSS selector-based content extraction
- 🚀 **Concurrent processing** - Parallel scraping for multiple sources
- 🌐 **Clean web interface** - Simple UI for managing scraping tasks
- ⚙️ **Configurable settings** - Customizable scraping parameters per source
- 🐳 **Docker support** - Containerized deployment with Docker Compose

## Project Structure

```
go-web-scraper/
├── cmd/
│   └── api/
│       └── main.go            # Application entry point
├── internal/
│   ├── api/
│   │   ├── handler/
│   │   │   ├── health.go      # Health check endpoint
│   │   │   └── scrape.go      # Scraping endpoint with SSE
│   │   └── router/
│   │       └── router.go      # HTTP router configuration
│   ├── scraper/
│   │   ├── core/
│   │   │   └── core.go        # Core scraping orchestration
│   │   └── parser/
│   │       └── parser.go      # HTML parsing and data extraction
│   └── shared/
│       ├── config/
│       │   └── config.go      # Configuration management
│       └── model/
│           ├── data.go        # Data structures for scraped content
│           └── types.go       # Request/response types
├── static/
│   └── index.html             # Web interface
├── docker-compose.yml         # Docker Compose configuration
├── go.mod                     # Go module dependencies
└── README.md                  # This file
```

## Dependencies

- **[Gin](https://github.com/gin-gonic/gin)** - HTTP web framework
- **[Colly](https://github.com/gocolly/colly)** - Web scraping framework
- **Frontend** - Vanilla HTML/CSS/JavaScript with streaming support

## Installation

### Local Development

1. **Clone the repository:**
   ```bash
   git clone https://github.com/OvidiuGi/go-web-scraper.git
   cd go-web-scraper
   ```

2. **Install dependencies:**
   ```bash
   go mod download
   ```

3. **Run the application:**
   ```bash
   go run cmd/api/main.go
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

The scraper includes specialized configurations for:

- **Euronews Romania** - `https://www.euronews.ro/categorii/romania`
  - Uses specific `/articole` path filtering
  - Extracts content from `div[itemprop='articleBody'] p` elements
  
- **Digi24** - `https://www.digi24.ro/stiri/actualitate/politica`
  - Flexible CSS selector-based extraction
  - Content parsing from `div.data-app-meta-article` containers
  - Support for lists and structured content

Additional sites can be added by configuring appropriate CSS selectors in the request payload.

### Custom Configurations

```go
type ScraperSettings struct {
    Source          string        `json:"source"`
    SourceSearchTag string        `json:"source_search_tag"`
    VisitChild      bool          `json:"visit_child"`
    ChildSettings   ChildSettings `json:"child_settings"`
}

type ChildSettings struct {
    SearchAttr string `json:"search_attr"`
    TitleAttr  string `json:"title_attr"`
}
```

## Architecture

### Core Components

- **`cmd/api/main.go`** - Application entry point with server initialization
- **`internal/api/handler/`** - HTTP request handlers for REST endpoints
- **`internal/api/router/`** - HTTP routing configuration
- **`internal/scraper/core/`** - Core scraping orchestration logic
- **`internal/scraper/parser/`** - HTML parsing and content extraction
- **`internal/shared/config/`** - Configuration management
- **`internal/shared/model/`** - Data models and type definitions

### Data Flow

1. **HTTP Request** → API Handler
2. **Configuration** → Scraper Core
3. **Colly Collectors** → Parser Functions
4. **Extracted Data** → SSE Stream
5. **Real-time Updates** → Web Interface

### Scraping Process

1. **Main Collector** visits the source URL
2. **CSS Selectors** find article links
3. **Child Collector** visits individual articles
4. **Parser Functions** extract title and content
5. **Data Structures** store results
6. **SSE Stream** sends real-time updates

## License

This project is open source and available under the [MIT License](LICENSE).

## Disclaimer

This tool is for educational and research purposes. Please respect:
- Website terms of service
- Copyright and content ownership

Implement appropriate delays and be respectful to target websites to avoid overloading their servers.
