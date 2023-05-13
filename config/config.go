package config

import (
	"flag"
	"os"
	"strings"
)

type AppConfig struct {
	Host      string
	ResultURL string
}

func GetCfg() *AppConfig {

	a := &AppConfig{}

	flag.StringVar(&a.Host, "a", "localhost:8080", "It's a Host")
	flag.StringVar(&a.ResultURL, "b", "http://localhost:8080", "It's a Result URL")

	if baseURL := os.Getenv("BASE_URL"); baseURL != "" {
		a.ResultURL = strings.TrimSpace(baseURL)
	}
	if host := os.Getenv("SERVER_ADDRESS"); host != "" {
		a.Host = strings.TrimSpace(host)
	}
	return a
}
