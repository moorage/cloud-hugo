package config

import (
	"encoding/json"
	"errors"
	"os"

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
	BaseDir     string `default:"./repo"`
	HostingDir  string `default:"./www"`
	AccessToken string `json:"access_token"`
	UserName    string `json:"user_name"`
	SubPort     string `default:"8081"`
}

// PublisherConfig is the config for the publisher
type PublisherConfig struct {
	*CommonConfig
	RepoURL string `json:"repo_url"`
	// reason for this default being website is because this is more user facing then the
	// subscriber
	BaseDir     string `default:"./website"`
	AccessToken string `json:"access_token"`
	UserName    string `json:"user_name"`
	UserEmail   string `json:"user_email"`
	PubPort     string `default:"8080"`
}

// NewSubsciberConfig creates a Config struct populating the Config with env variables having prefix
// "CHUGO_SUB"
func NewSubsciberConfig() (*SubscriberConfig, error) {
	var subConf SubscriberConfig
	err := LoadFromFile("config/sub-config.json", &subConf)

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
	err := LoadFromFile("config/pub-config.json", &pubConf)
	if err != nil {
		return nil, err
	}
	err = envconfig.Process("CHUGO_PUB", &pubConf)
	if err != nil {
		return nil, err
	}

	if err = pubConf.Validate(); err != nil {
		return nil, err
	}

	return &pubConf, nil
}

// Validate validates a publisher config
func (pcg *PublisherConfig) Validate() error {
	// TODO: use reflect maybe?
	if pcg.RepoURL == "" {
		return errors.New("Repository URL must be provided in the config")
	}
	if pcg.UserName == "" {
		return errors.New("Username cannot be blank in the config")
	}
	if pcg.UserEmail == "" {
		return errors.New("User Email cannot be blank in the config")
	}
	return nil
}

// LoadFromFile gets the config from a file
func LoadFromFile(filename string, config interface{}) error {
	configFile, err := os.Open(filename)
	defer configFile.Close()
	if err != nil {
		return err
	}
	jsonParser := json.NewDecoder(configFile)
	return jsonParser.Decode(config)
}
