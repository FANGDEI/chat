package cacher

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"log"
)

func init() {
	var err error
	defaultCacherManager, err = New()
	if err != nil {
		log.Println(err)
	}
}

var defaultCacherManager *Manager

type Manager struct {
	handler *redis.Client
}

func New() (*Manager, error) {
	m := &Manager{
		handler: redis.NewClient(
			&redis.Options{
				Addr:     C.Addr + ":" + C.Port,
				Password: C.Password,
			},
		),
	}
	return m, m.handler.Ping(context.Background()).Err()
}

func GetDefaultCacherManager() *Manager {
	return defaultCacherManager
}

func (m *Manager) getEmailVerifyKey(email string) string {
	return fmt.Sprintf(
		"email:code:%s",
		email,
	)
}

func (m *Manager) getEmailBanKey(email string) string {
	return fmt.Sprintf(
		"email:ban:%s",
		email,
	)
}
