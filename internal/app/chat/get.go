package chat

import (
	"chat/internal/app/service"
	"chat/internal/pkg/errno"
	"context"
	"time"
)

func (m *Manager) Get(ctx context.Context, request *service.GetRequest) (*service.GetResponse, error) {
	m.logger.Info("Chat Service, Get service")
	var expiration time.Duration
	var response service.GetResponse
	ddl, ok := ctx.Deadline()
	if !ok {
		expiration = time.Hour
	} else {
		expiration = ddl.Sub(time.Now())
	}
	msgs, err := m.cacher.GetMsg(request.Uuid, expiration)
	if err != nil {
		return nil, errno.ServerErr(errno.ErrGetUserMsg, err.Error())
	}
	select {
	case <-ctx.Done():
		if err := m.cacher.Rewrite(request.Uuid, msgs); err != nil {
			return nil, errno.ServerErr(errno.ErrRewriteMsg, err.Error())
		}
	default:
		if err := m.cacher.CreateHistory(request.Uuid, msgs); err != nil {
			return nil, errno.ServerErr(errno.ErrCreateHistory, err.Error())
		}
		response.Msg = msgs
	}
	return &response, nil
}
