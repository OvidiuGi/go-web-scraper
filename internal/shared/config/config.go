package config

type Config struct {
	ApiPort string `env:"PORT" default:"8080"`
}

func LoadConfig() (*Config, error) {
	// Load configuration from environment variables or a config file
	// For simplicity, we'll return a hardcoded config here
	return &Config{
		ApiPort: "8080",
	}, nil
}
