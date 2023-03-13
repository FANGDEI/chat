package cacher

import (
	"context"
	"encoding/json"
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
