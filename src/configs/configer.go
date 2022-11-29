package configs

import (
	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	DevUrl  string `envconfig:"DEV_URL" default:"http:localhost:8080"`
	DevPort string `envconfig:"DEV_PORT" default:"8080"`
}

func New() (*Config, error) {

	conf := &Config{}
	err := envconfig.Process("", conf)
	if err != nil {
		return nil, err
	}

	return conf, nil

}
