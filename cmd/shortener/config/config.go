package config

import (
	"encoding/json"
	"flag"
	"log"
	"os"

	"github.com/caarlos0/env"
)

const (
	defaultServerAddress = "localhost:8080"
	defaultBaseURL       = "http://localhost:8080"
)

// AppConfig - Configuration
type AppConfig struct {
	Host           string `env:"SERVER_ADDRESS" json:"server_address"`
	ResultURL      string `env:"BASE_URL" json:"base_url"`
	FilePATH       string `env:"FILE_STORAGE_PATH" json:"file_storage_path"`
	DataBaseString string `env:"DATABASE_DSN" json:"database_dsn"`
	EnableHTTPS    bool   `env:"ENABLE_HTTPS" json:"enable_https"`
	ConfigJSON     string `env:"CONFIG" json:"-"`
	TrustedSubnet  string `env:"TRUSTED_SUBNET" json:"trusted_subnet"`
}

func (a *AppConfig) loadConfigFromFile(file string) {

	data, err := os.ReadFile(file)
	if err != nil {
		log.Println(err)
	}

	err = json.Unmarshal(data, &a)
	if err != nil {
		log.Println(err, "Error parsing config JSON")
	}

}

// NewAppConfig - AppConfig constructor
func NewAppConfig() *AppConfig {

	a := AppConfig{}

	flag.StringVar(&a.Host, "a", defaultServerAddress, "It's a Host")
	flag.StringVar(&a.ResultURL, "b", defaultBaseURL, "It's a Result URL")
	flag.StringVar(&a.FilePATH, "f", "", "It's a FilePATH")
	flag.StringVar(&a.DataBaseString, "d", "", "it's conn string")
	flag.BoolVar(&a.EnableHTTPS, "s", false, "using HTTPS")
	flag.StringVar(&a.ConfigJSON, "c", "", "It's a ConfigJSON file")
	flag.StringVar(&a.TrustedSubnet, "t", "", "It's a Trusted Subnet")

	flag.Parse()

	err := env.Parse(&a)
	if err != nil {
		panic(err)
	}

	if a.ConfigJSON != "" {

		j := AppConfig{}
		j.loadConfigFromFile(a.ConfigJSON)

		if j.Host != defaultServerAddress && j.Host != "" {
			a.Host = j.Host
		}
		if a.ResultURL != defaultServerAddress && j.ResultURL != "" {
			a.ResultURL = j.ResultURL
		}
		if a.DataBaseString == "" && j.DataBaseString != "" {
			a.DataBaseString = j.DataBaseString
		}
		if a.FilePATH == "" && j.FilePATH != "" {
			a.FilePATH = j.FilePATH
		}
		if j.EnableHTTPS {
			a.EnableHTTPS = true
		}
		if a.TrustedSubnet == "" && j.TrustedSubnet != "" {
			a.TrustedSubnet = j.TrustedSubnet
		}

	}

	return &a
}
