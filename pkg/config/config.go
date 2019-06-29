package config

import (
	"github.com/kelseyhightower/envconfig"
)

// CommonConfig is the global config used through out the application, both the publisher and subscriber
type CommonConfig struct {
	CredFile  string `default:"credentials.json"`
	ProjectID string `default:"cloud-hugo-test"`
	TopicName string `default:"chugo-run-requests"`
	Env       string `default:"dev"`
}

// SubscriberConfig is the config for the subscriber
type SubscriberConfig struct {
	*CommonConfig
	BaseDir string `default:"./repos"`
}

// NewSubsciber creates a Config struct populating the Config folder with env variables having prefix
// "CHUGO_"
func NewSubsciber() *SubscriberConfig {
	var c SubscriberConfig
	err := envconfig.Process("CHUGO", &c)

	if err != nil {
		panic(err)
	}
	return &c
}
