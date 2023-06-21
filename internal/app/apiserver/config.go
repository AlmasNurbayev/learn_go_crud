package apiserver

import "github.com/AlmasNurbayev/learn_go_crud/internal/app/store"

type Config struct {
	BindAddr string
	LogLevel string
	Store    *store.Config
}

func NewConfig() *Config {
	return &Config{
		BindAddr: "8083",
		LogLevel: "debug",
		Store:    store.NewConfig(),
	}
}
