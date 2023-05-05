package config

import (
	"flag"
	"os"
	"strings"
)

var AppConfig struct {
	Host      string
	ResultURL string
}

//инициализируем флаги и переменные окружения

func init() {

	if host := os.Getenv("SERVER_ADDRESS"); host != "" {
		AppConfig.Host = strings.TrimSpace(host)
	} else {
		flag.StringVar(&AppConfig.Host, "a", "localhost:8080", "It's a Host")
	}
	if baseURL := os.Getenv("BASE_URL"); baseURL != "" {
		AppConfig.ResultURL = strings.TrimSpace(baseURL)
	} else {
		flag.StringVar(&AppConfig.ResultURL, "b", "http://localhost:8080", "It's a Result URL")
	}

	//flag.StringVar(&AppConfig.Host, "a", "localhost:8080", "It's a Host")
	//flag.StringVar(&AppConfig.ResultURL, "b", "http://localhost:8080", "It's a Result URL")

}
