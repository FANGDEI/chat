package cacher

import (
	"context"
	"fmt"
	"log"

	"github.com/go-redis/redis/v8"
)

func init() {
	var err error
	defaultCacherManager, err = New()
	if err != nil {
		log.Println(err)
	}
}

var defaultCacherManager *Manager

type Message struct {
	From        int64  `json:"from"`
	To          int64  `json:"to"`
	Content     string `json:"content"`
	Time        string `json:"time"`
	ContentType int64  `json:"content_type"`
	MessageType int64  `json:"message_type"`
}

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

func (m *Manager) getMsgReceiverKey(id int64) string {
	return fmt.Sprintf(
		"chat:msg:%d", id,
	)
}

func (m *Manager) getHistoryKey(id, otherID int64) string {
	return fmt.Sprintf(
		"history:%d-%d", id, otherID,
	)
}
