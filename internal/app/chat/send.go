package chat

import (
	"chat/internal/app/service"
	"chat/internal/pkg/cacher"
	"chat/internal/pkg/constanter"
	"chat/internal/pkg/errno"
	"context"
	"time"
)

func (m *Manager) Send(ctx context.Context, request *service.SendRequest) (*service.Response, error) {
	m.logger.Info("Chat Service, Send Service")
	// 私聊
	if request.MessageType == constanter.MESSAGE_TYPE_USER {
		data := &cacher.Message{
			From:        request.From,
			To:          request.To,
			GroupID:     request.GroupId,
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
	// 群聊
	list, err := m.localer.GetGroupMembersIDWithGroupID(request.To)
	if err != nil {
		return nil, errno.ServerErr(errno.ErrDatabase, err.Error())
	}
	for _, id := range list {
		// 群聊遇到自己跳过
		if id == request.From {
			continue
		}
		err := m.cacher.Send(&cacher.Message{
			From:        request.From,
			To:          id,
			GroupID:     request.To,
			Content:     request.Content,
			ContentType: request.ContentType,
			MessageType: request.MessageType,
			Time:        time.Now().Format("2006.01.02 15:04:05"),
		})
		if err != nil {
			return nil, errno.ServerErr(errno.ErrRedis, err.Error())
		}
	}
	return &service.Response{}, nil
}
