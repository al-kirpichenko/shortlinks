package config

import (
	"flag"
	"fmt"
	"github.com/caarlos0/env"
)

const (
	DBhost     = "localhost"
	DBuser     = "postgres"
	DBpassword = ""
	DBdbname   = "shorten"
)

type AppConfig struct {
	Host           string `env:"SERVER_ADDRESS"`
	ResultURL      string `env:"BASE_URL"`
	FilePATH       string `env:"FILE_STORAGE_PATH"`
	DataBaseString string `env:"DATABASE_DSN"`
}

func NewAppConfig() *AppConfig {

	a := AppConfig{}

	ps := fmt.Sprintf("host=%s user=%s password=%s dbname=%s sslmode=disable",
		DBhost, DBuser, DBpassword, DBdbname)

	flag.StringVar(&a.Host, "a", "localhost:8080", "It's a Host")
	flag.StringVar(&a.ResultURL, "b", "http://localhost:8080", "It's a Result URL")
	flag.StringVar(&a.FilePATH, "f", "/tmp/short-url-db.json", "It's a FilePATH")
	flag.StringVar(&a.DataBaseString, "d", ps, "it's conn string")

	flag.Parse()

	err := env.Parse(&a)
	if err != nil {
		panic(err)
	}

	return &a
}
