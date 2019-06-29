package config

import (
	"encoding/json"
	"os"

	"github.com/kelseyhightower/envconfig"
)

// CommonConfig is the global config used through out the application, both the publisher and subscriber
type CommonConfig struct {
	CredFile    string `default:"credentials.json"`
	ProjectID   string `default:"cloud-hugo-test"`
	TopicName   string `default:"chugo-run-requests"`
	Env         string `default:"dev"`
	AccessToken string
}

// SubscriberConfig is the config for the subscriber
type SubscriberConfig struct {
	*CommonConfig
	BaseDir string `default:"./repos"`
}

type PublisherConfig struct {
	*CommonConfig
}

// NewSubsciber creates a Config struct populating the Config folder with env variables having prefix
// "CHUGO_"
func NewSubsciber() (*SubscriberConfig, error) {
	var subConf SubscriberConfig
	err := LoadFromFile("sub-config.json", &subConf)

	if err != nil {
		return nil, err
	}
	err = envconfig.Process("CHUGO", &subConf)
	if err != nil {
		return nil, err
	}

	return &subConf, nil
}

// LoadFromFile gets the config from a file
func LoadFromFile(filename string, config interface{}) error {
	configFile, err := os.Open(filename)
	defer configFile.Close()
	if err != nil {
		return err
	}
	jsonParser := json.NewDecoder(configFile)
	return jsonParser.Decode(&config)
}
