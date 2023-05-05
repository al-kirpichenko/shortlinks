package config

import "flag"

var AppConfig struct {
	Host      string
	ResultURL string
}

func init() {

	//AppConfig.Host = "localhost:8080"
	//AppConfig.ResultURL = "http://localhost:8080"

	flag.StringVar(&AppConfig.Host, "a", "localhost:8080", "It's a Host")
	flag.StringVar(&AppConfig.ResultURL, "b", "http://localhost:8080", "It's a Result URL")

}
