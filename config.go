package latrappemelder

import (
	"fmt"
	"os"
	"strconv"
)

var (
	defaultHTTPAddress    = ":1234"
	defaultMetricsAddress = ":9090"
	defaultDatabaseURL    = "sqlite:test.db"
	defaultLaTrappeURL    = "https://www.kloosterwinkelonline.nl/la-trappe-trappist-oak-aged"
)

// Config contains all the config for serving la trappe melder
type Config struct {
	AppURL         string
	HTTPAddress    string
	MetricsAddress string
	AccessLog      bool
	DatabaseURL    string
	LaTrappeURL    string

	SMTP struct {
		Host       string
		Port       int
		User       string
		Password   string
		DisableTLS bool
		FromEmail  string
		FromName   string
	}

	AdminMail string
}

// BuildConfigFromEnv populates a config from env variables
func BuildConfigFromEnv() *Config {
	config := &Config{}

	config.AppURL = getEnv("APP_URL", "")
	config.HTTPAddress = getEnv("HTTP_ADDRESS", defaultHTTPAddress)
	config.MetricsAddress = getEnv("METRICS_ADDRESS", defaultMetricsAddress)
	config.DatabaseURL = getEnv("DATABASE_URL", defaultDatabaseURL)
	config.LaTrappeURL = getEnv("LATRAPPE_URL", defaultLaTrappeURL)

	config.SMTP.Host = getEnv("SMTP_HOST", "")

	port, err := strconv.Atoi(getEnv("SMTP_PORT", "0"))
	if err != nil {
		config.SMTP.Port = 0
	} else {
		config.SMTP.Port = port
	}

	config.SMTP.User = getEnv("SMTP_USER", "")
	config.SMTP.Password = getEnv("SMTP_PASSWORD", "")
	if getEnv("SMTP_DISABLE_TLS", "0") == "1" {
		config.SMTP.DisableTLS = true
	}

	config.SMTP.FromName = getEnv("SMTP_FROM_NAME", "")
	config.SMTP.FromEmail = getEnv("SMTP_FROM_EMAIL", "")

	// access log
	accessLog := getEnv("ACCESS_LOG", "1")
	if accessLog == "0" {
		config.AccessLog = false
	}

	config.AdminMail = getEnv("ADMIN_MAIL", "")

	return config
}

// Validate validates whether all config is set and valid
func (config *Config) Validate() error {

	if config.AppURL == "" {
		return fmt.Errorf("APP_URL cannot be empty")
	}
	if config.HTTPAddress == "" {
		return fmt.Errorf("HTTP_ADDRESS cannot be empty")
	}
	if config.MetricsAddress == "" {
		return fmt.Errorf("METRICS_ADDRESS cannot be empty")
	}

	if config.SMTP.Host == "" {
		return fmt.Errorf("SMTP_HOST must be set")
	}
	if config.SMTP.Port == 0 {
		return fmt.Errorf("SMTP_PORT must be set")
	}
	if config.SMTP.FromEmail == "" {
		return fmt.Errorf("SMTP_FROM_EMAIL must be set")
	}

	if config.AdminMail == "" {
		return fmt.Errorf("ADMIN_MAIL must be set")
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
