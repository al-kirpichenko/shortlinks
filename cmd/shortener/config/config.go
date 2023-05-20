package config

import (
	"flag"
	"github.com/caarlos0/env/v6"
	"log"
)

type AppConfig struct {
	Host      string `env:"SERVER_ADDRESS"`
	ResultURL string `env:"BASE_URL"`
	FilePATH  string `env:"FILE_STORAGE_PATH"`
}

func NewCfg() *AppConfig {

	a := &AppConfig{}

	flag.StringVar(&a.Host, "a", "localhost:8080", "It's a Host")
	flag.StringVar(&a.ResultURL, "b", "http://localhost:8080", "It's a Result URL")
	flag.StringVar(&a.FilePATH, "f", "/tmp/short-url-db.json", "It's a FilePATH")

	err := env.Parse(a)
	if err != nil {
		log.Fatal(err)
	}

	return a
}
