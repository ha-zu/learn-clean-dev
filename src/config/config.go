package config

import (
	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	DevURL  string `envconfig:"DEV_URL" default:"http:localhost:8080"`
	DevPort string `envconfig:"DEV_PORT" default:"8080"`
}

var Env Config

func init() {
	EnvSetting()
	err := envconfig.Process("", &Env)
	if err != nil {
		// ToDo: Error Handling
		panic(err)
	}
}
