package cacher

import (
	"context"
	"encoding/json"
	"time"
)

// Send received message to Redis List
// Use the structure of Redis List as message queue
// To push instant messages and store offline messages
func (m *Manager) Send(message *Message) error {
	msg, err := json.Marshal(message)
	if err != nil {
		return err
	}
	return m.handler.LPush(context.Background(), m.getMsgReceiverKey(message.To), msg).Err()
}

func (m *Manager) GetMsg(uuid string, expiration time.Duration) ([]string, error) {
	key := m.getMsgReceiverKey(uuid)
	if m.handler.LLen(context.Background(), key).Val() == 0 {
		// 阻塞读
		msg, err := m.handler.BRPop(context.Background(), expiration, key).Result()
		if err != nil {
			return nil, err
		}
		// msg[0] = chat:msg:key
		return msg[1:], nil
	}
	var msgs []string
	for m.handler.LLen(context.Background(), key).Val() != 0 {
		msg, err := m.handler.RPop(context.Background(), key).Result()
		if err != nil {
			return nil, err
		}
		msgs = append(msgs, msg)
	}
	return msgs, nil
}

// Rewrite
func (m *Manager) Rewrite(uuid string, msgs []string) error {
	key := m.getMsgReceiverKey(uuid)
	for i := len(msgs) - 1; i >= 0; i-- {
		if err := m.handler.RPush(context.Background(), key, msgs[i]).Err(); err != nil {
			return err
		}
	}
	return nil
}

// CreateHistory
func (m *Manager) CreateHistory(uuid string, data []string) error {
	return nil
}
