package main

import "github.com/Netflix/go-env"

type Config struct {
	APIKey     string `env:"API_KEY,required=true"`
	ListenPort int    `env:"LISTEN_PORT,default=49153"`
}

func NewConfig() (*Config, error) {
	var config Config

	_, err := env.UnmarshalFromEnviron(&config)
	if err != nil {
		return nil, err
	}

	return &config, nil
}
