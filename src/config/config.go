package config

import (
	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	DevURL  string `envconfig:"DEV_URL" default:"http:localhost:1323"`
	DevPort string `envconfig:"DEV_PORT" default:"1323"`
}

var Env Config

func init() {
	err := envconfig.Process("", &Env)
	if err != nil {
		// ToDo: Error Handling
		panic(err)
	}
}
