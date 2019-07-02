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
	UserName    string
}

// SubscriberConfig is the config for the subscriber
type SubscriberConfig struct {
	*CommonConfig
	BaseDir string `default:"./repos"`
}

// PublisherConfig is the config for the publisher
type PublisherConfig struct {
	*CommonConfig
	RepoURL string `json:"repo_url"`
}

// NewSubsciberConfig creates a Config struct populating the Config with env variables having prefix
// "CHUGO_SUB"
func NewSubsciberConfig() (*SubscriberConfig, error) {
	var subConf SubscriberConfig
	err := LoadFromFile("sub-config.json", &subConf)

	if err != nil {
		return nil, err
	}
	err = envconfig.Process("CHUGO_SUB", &subConf)
	if err != nil {
		return nil, err
	}

	return &subConf, nil
}

// NewPublisherConfig creates a Config struct populating the Config with env variables having prefix
// "CHUGO_PUB"
func NewPublisherConfig() (*PublisherConfig, error) {
	var pubConf PublisherConfig
	err := LoadFromFile("pub-config.json", &pubConf)

	if err != nil {
		return nil, err
	}
	err = envconfig.Process("CHUGO_PUB", &pubConf)
	if err != nil {
		return nil, err
	}

	return &pubConf, nil
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
