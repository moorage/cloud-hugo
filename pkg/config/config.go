package config

import (
	"github.com/kelseyhightower/envconfig"
)

// Config is the global config used through out the application
type Config struct {
	CredFile  string `default:"credentials.json"`
	ProjectID string `default:"cloud-hugo-test"`
	TopicName string `default:"chugo-run-requests"`
	Env       string `default:"dev"`
}

// New creates a Config struct populating the Config folder with env variables having prefix
// "CHUGO_"
func New() *Config {
	var c Config
	err := envconfig.Process("CHUGO", &c)

	if err != nil {
		panic(err)
	}
	return &c
}
