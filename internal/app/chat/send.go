package chat

import (
	"chat/internal/app/service"
	"chat/internal/pkg/cacher"
	"chat/internal/pkg/errno"
	"context"
	"time"
)

func (m *Manager) Send(ctx context.Context, request *service.SendRequest) (*service.Response, error) {
	m.logger.Info("Chat Service, Send Service")
	data := &cacher.Message{
		From:        request.From,
		To:          request.To,
		Content:     request.Content,
		ContentType: request.ContentType,
		MessageType: request.MessageType,
		Time:        time.Now().Format("2006.01.02 15:04:05"),
	}
	err := m.cacher.Send(data)
	if err != nil {
		return nil, errno.ServerErr(errno.ErrRedis, err.Error())
	}
	return &service.Response{}, nil
}
