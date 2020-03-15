package config

import (
	"fmt"
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

// Load returns Configuration struct
func Load(env string) (*Configuration, error) {
	bytes, err := ioutil.ReadFile("config." + env + ".yaml")
	if err != nil {
		return nil, fmt.Errorf("error reading config file, %s", err)
	}
	var cfg = new(Configuration)
	if err := yaml.Unmarshal(bytes, cfg); err != nil {
		return nil, fmt.Errorf("unable to decode into struct, %v", err)
	}
	return cfg, nil
}

// Configuration holds data necessery for configuring application
type Configuration struct {
	JWT    *JWT      `yaml:"jwt"`
	DB     *Database `yaml:"db"`
	Casbin *Casbin   `yaml:"casbin"`
}

// App holds data for internal app configuration
type App struct {
	ResetPasswordLinkExpire   int     `yaml:"reset_password_link_expire"`
	SendgridAPIKey            string  `yaml:"sendgrid_api_key"`
	SentryDSN                 string  `yaml:"sentry_dsn"`
	CompressUploadedImage     bool    `yaml:"compress_uploaded_image"`
	CompressImageMinThreshold float64 `yaml:"compress_uploaded_image_min_threshold"`
}

// JWT holds data necessery for JWT configuration
type JWT struct {
	Secret           string
	Duration         int // in hours
	SigningAlgorithm string
}

// Database holds data necessery for database configuration
type Database struct {
	PSN          string
	Log          bool
	CreateSchema bool
	Timeout      int
}

// Casbin holds data casbin configuration
type Casbin struct {
	Model  string
	Policy string
}
