package logger

import (
	"log"

	config "chat/internal/pkg/configer"
)

type Config struct {
	Output      string `json:"output"`
	ProjectName string `json:"project_name"`
	Level       string `json:"level"`
}

func (c *Config) Key() string {
	return "chat/logger"
}

var C Config

func init() {
	err := config.ReadConfig(&C)
	if err != nil {
		log.Println(err)
	}
}
