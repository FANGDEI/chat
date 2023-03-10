package web

import (
	config "chat/internal/pkg/configer"
	"log"
)

type Config struct {
	Host string `json:"host"`
	Port string `json:"port"`
}

func (c *Config) Key() string {
	return "chat/api"
}

var C Config

func init() {
	err := config.ReadConfig(&C)
	if err != nil {
		log.Fatalln(err)
	}
}
