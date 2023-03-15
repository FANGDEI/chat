package obser

import (
	"chat/internal/pkg/configer"
	"log"
)

type Config struct {
	Address   string `json:"address"`
	SecretID  string `json:"secret_id"`
	SecretKey string `json:"secret_key"`
	Bucket    string `json:"bucket"`
}

func (c *Config) Key() string {
	return "chat/obs"
}

var C Config

func init() {
	err := configer.ReadConfig(&C)
	if err != nil {
		log.Fatalf("failed to load config %v, errno: %v", C.Key(), err)
	}
}
