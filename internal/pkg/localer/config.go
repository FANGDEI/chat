package localer

import (
	"log"

	config "chat/internal/pkg/configer"
)

type Config struct {
	Host     string `json:"host"`
	Port     string `json:"port"`
	User     string `json:"user"`
	Name     string `json:"name"`
	Password string `json:"password"`
}

func (c *Config) Key() string {
	return "chat/local"
}

var C Config

func init() {
	err := config.ReadConfig(&C)
	if err != nil {
		log.Println(err)
	}
}
