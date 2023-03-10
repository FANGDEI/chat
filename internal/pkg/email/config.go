package email

import (
	config "chat/internal/pkg/configer"
	"log"
)

type Config struct {
	Account  string `json:"account"`
	Name     string `json:"name"`
	Password string `json:"password"`
	Service  string `json:"service"`
	Port     int    `json:"port"`
}

func (c *Config) Key() string {
	return "chat/email"
}

var C Config

func init() {
	err := config.ReadConfig(&C)
	if err != nil {
		log.Fatalln(err)
	}
	defaultEmailerManager = New()
	go defaultEmailerManager.Run()
}
