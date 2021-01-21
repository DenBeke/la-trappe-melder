package latrappemelder

import (
	"fmt"
	"os"
)

var (
	defaultHTTPAddress    = ":1234"
	defaultMetricsAddress = ":9090"
	defaultDatabaseURL    = "sqlite:test.db"
	defaultLaTrappeURL    = "https://www.kloosterwinkelonline.nl/la-trappe-trappist-oak-aged"
)

// Config contains all the config for serving la trappe melder
type Config struct {
	HTTPAddress    string
	MetricsAddress string
	AccessLog      bool
	DatabaseURL    string
	LaTrappeURL    string
}

// BuildConfigFromEnv populates a config from env variables
func BuildConfigFromEnv() *Config {
	config := &Config{}

	config.HTTPAddress = getEnv("HTTP_ADDRESS", defaultHTTPAddress)
	config.MetricsAddress = getEnv("HTTP_ADDRESS", defaultMetricsAddress)
	config.DatabaseURL = getEnv("DATABASE_URL", defaultDatabaseURL)
	config.LaTrappeURL = getEnv("DATABASE_URL", defaultLaTrappeURL)

	// access log
	accessLog := getEnv("ACCESS_LOG", "1")
	if accessLog == "0" {
		config.AccessLog = false
	}

	return config
}

// Validate validates whether all config is set and valid
func (config *Config) Validate() error {

	if config.HTTPAddress == "" {
		return fmt.Errorf("HTTPAddress cannot be empty")
	}
	if config.MetricsAddress == "" {
		return fmt.Errorf("MetricsAddress cannot be empty")
	}

	return nil
}

// getEnv gets the env variable with the given key if the key exists
// else it falls back to the fallback value
func getEnv(key, fallback string) string {
	value := os.Getenv(key)
	if len(value) == 0 {
		return fallback
	}
	return value
}
