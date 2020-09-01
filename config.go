package epcc

import (
	"log"
	"time"

	"github.com/kelseyhightower/envconfig"
)

// cfg holds the current configuration.
var cfg Config

// initialisation for the epcc package
func init() {

	// Set default configuration values. 
	cfg.BaseURL = "https://api.moltin.com/"
	cfg.ClientTimeout = 10 * time.Second
	cfg.RetryLimitTimeout = 30 * time.Second

	// process environment variables and store them in the global cfg.
	err := envconfig.Process("", &cfg)
	if err != nil {
		log.Fatal(err.Error())
	}

	// is the next line needed will a missing env var catch this?
	if cfg.Credentials.ClientID == "" {
		log.Fatal("Required environment variable GO_EPCC_CLIENT_ID not found")
	}
	if cfg.Credentials.ClientSecret == "" {
		log.Fatal("Required environment variable GO_EPCC_CLIENT_SECRET not found")
	}
}

// Config is used to keep track of configuration in one place.
// fields tagged envconfig are read from environment variables.
// fields tagged default are default values.
type Config struct {
	Credentials struct {
		ClientID     string `envconfig:"GO_EPCC_CLIENT_ID"`
		ClientSecret string `envconfig:"GO_EPCC_CLIENT_SECRET"`
	}
		BaseURL string
		ClientTimeout time.Duration
		RetryLimitTimeout time.Duration
}
