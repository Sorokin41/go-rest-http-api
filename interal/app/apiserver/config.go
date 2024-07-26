package apiserver

import "github.com/Sorokin41/go-rest-http-api/store"

type Config struct {
	BindAddr string        `toml:"bind_addr"`
	Loglevel string        `toml:"log_level"`
	Store    *store.Config `toml:"store"`
}

func NewConfig() *Config {
	return &Config{
		BindAddr: ":8080",
		Loglevel: "debug",
		Store:    store.NewConfig(),
	}
}
