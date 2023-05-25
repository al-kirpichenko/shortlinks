package config

import (
	"flag"
	//"os"
	//"strings"
	"github.com/caarlos0/env"
)

type AppConfig struct {
	Host      string `env:"SERVER_ADDRESS"`
	ResultURL string `env:"BASE_URL"`
	FilePATH  string `env:"FILE_STORAGE_PATH"`
}

func NewAppConfig() *AppConfig {

	a := AppConfig{}

	flag.StringVar(&a.Host, "a", "localhost:8080", "It's a Host")
	flag.StringVar(&a.ResultURL, "b", "http://localhost:8080", "It's a Result URL")
	flag.StringVar(&a.FilePATH, "f", "/tmp/short-url-db.json", "It's a FilePATH")
	flag.Parse()

	err := env.Parse(&a)
	if err != nil {
		panic(err)
	}
	//if baseURL := os.Getenv("BASE_URL"); baseURL != "" {
	//	a.ResultURL = strings.TrimSpace(baseURL)
	//}
	//if host := os.Getenv("SERVER_ADDRESS"); host != "" {
	//	a.Host = strings.TrimSpace(host)
	//}
	//if file := os.Getenv("FILE_STORAGE_PATH"); file != "" {
	//	a.FilePATH = file
	//}

	return &a
}
