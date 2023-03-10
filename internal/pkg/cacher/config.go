package cacher

import (
	config "chat/internal/pkg/configer"
	"log"
)

type Config struct {
	Addr     string `json:"addr"`
	Port     string `json:"port"`
	Password string `json:"password"`
}

func (c *Config) Key() string {
	return "chat/cache"
}

var C Config

func init() {
	err := config.ReadConfig(&C)
	if err != nil {
		log.Println(err)
	}
}
