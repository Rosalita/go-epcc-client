package epcc

import (
	"log"

	"github.com/kelseyhightower/envconfig"
)

// Cfg holds the current configuration.
var Cfg Config

// initialisation for the epcc package
func init() {
	// environment variables are processed using envconfig and stored in the global cfg.
	err := envconfig.Process("", &Cfg)
	if err != nil {
		log.Fatal(err.Error())
	}

	if Cfg.Credentials.ClientID == "" {
		log.Fatal("Required environment variable GO_EPCC_CLIENT_ID not found")
	}
	if Cfg.Credentials.ClientSecret == "" {
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
	BaseURL string `default:"https://api.moltin.com/"`
}
