package config

import (
	"flag"
	"os"
	"strings"
)

type AppConfig struct {
	Host      string
	ResultURL string
	FilePATH  string
}

func flags(cfg *AppConfig) *AppConfig {

	flag.StringVar(&cfg.Host, "a", "localhost:8080", "It's a Host")
	flag.StringVar(&cfg.ResultURL, "b", "http://localhost:8080", "It's a Result URL")
	flag.StringVar(&cfg.FilePATH, "f", "/tmp/short-url-db.json", "It's a FilePATH")

	return cfg
}

func paths(cfg *AppConfig) *AppConfig {
	if baseURL := os.Getenv("BASE_URL"); baseURL != "" {
		cfg.ResultURL = strings.TrimSpace(baseURL)
	}
	if host := os.Getenv("SERVER_ADDRESS"); host != "" {
		cfg.Host = strings.TrimSpace(host)
	}
	if file := os.Getenv("FILE_STORAGE_PATH"); file != "" {
		cfg.FilePATH = file
	}
	return cfg
}

func NewAppConfig() *AppConfig {

	a := &AppConfig{}

	//flag.StringVar(&a.Host, "a", "localhost:8080", "It's a Host")
	//flag.StringVar(&a.ResultURL, "b", "http://localhost:8080", "It's a Result URL")
	//flag.StringVar(&a.FilePATH, "f", "/tmp/short-url-db.json", "It's a FilePATH")
	//
	//flag.Parse()

	cfg := flags(a)
	cfg = paths(a)

	//if baseURL := os.Getenv("BASE_URL"); baseURL != "" {
	//	a.ResultURL = strings.TrimSpace(baseURL)
	//}
	//if host := os.Getenv("SERVER_ADDRESS"); host != "" {
	//	a.Host = strings.TrimSpace(host)
	//}
	//if file := os.Getenv("FILE_STORAGE_PATH"); file != "" {
	//	a.FilePATH = file
	//}
	return cfg
}
