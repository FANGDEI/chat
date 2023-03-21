package cacher

import (
	"chat/internal/pkg/constanter"
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

func (m *Manager) GetMsg(id int64, expiration time.Duration) ([]string, error) {
	key := m.getMsgReceiverKey(id)
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
func (m *Manager) Rewrite(id int64, msgs []string) error {
	key := m.getMsgReceiverKey(id)
	for i := len(msgs) - 1; i >= 0; i-- {
		if err := m.handler.RPush(context.Background(), key, msgs[i]).Err(); err != nil {
			return err
		}
	}
	return nil
}

// RewriteAndPop Rewrite msgs to redis and Pop the history from redis
func (m *Manager) RewriteAndPop(id int64, msgs []string) error {
	key := m.getMsgReceiverKey(id)
	for i := len(msgs) - 1; i >= 0; i-- {
		if err := m.handler.RPush(context.Background(), key, msgs[i]).Err(); err != nil {
			return err
		}
		var msg Message
		err := json.Unmarshal([]byte(msgs[i]), &msg)
		if err != nil {
			return err
		}
		if err := m.handler.LPop(context.Background(), m.getHistoryKey(id, msg.From)).Err(); err != nil {
			return err
		}
	}
	return nil
}

// CreateHistory
func (m *Manager) CreateHistory(id int64, msgs []string) error {
	for i := len(msgs) - 1; i >= 0; i-- {
		var msg Message
		if err := json.Unmarshal([]byte(msgs[i]), &msg); err != nil {
			return err
		}
		// 好友申请以及入群申请不存入聊天记录
		if msg.ContentType == constanter.FRIEND_REQUEST || msg.ContentType == constanter.GROUP_REQUEST {
			continue
		}
		// 群聊消息记录
		if msg.MessageType == constanter.MESSAGE_TYPE_GROUP {
			key := m.getGroupHistoryKey(msg.GroupID)
			if err := m.handler.LPush(context.Background(), key, msgs[i]).Err(); err != nil {
				return err
			}
			continue
		}
		key := m.getHistoryKey(id, msg.From)
		if err := m.handler.LPush(context.Background(), key, msgs[i]).Err(); err != nil {
			return err
		}
	}
	return nil
}

func (m *Manager) GetUserHistory(userID, otherID, offset, limit int64, pagination bool) ([]string, error) {
	key := m.getHistoryKey(userID, otherID)
	var start, end int64 = 0, -1
	if pagination {
		start, end = offset, offset+limit
	}
	list, err := m.handler.LRange(context.Background(), key, start, end).Result()
	if err != nil {
		return nil, err
	}
	return list, err
}

func (m *Manager) GetGroupHistory(groupID, offset, limit int64, pagination bool) ([]string, error) {
	key := m.getGroupHistoryKey(groupID)
	var start, end int64 = 0, -1
	if pagination {
		start, end = offset, offset+limit
	}
	list, err := m.handler.LRange(context.Background(), key, start, end).Result()
	if err != nil {
		return nil, err
	}
	return list, err
}
