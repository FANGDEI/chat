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
	msgs, err := m.cacher.GetMsg(request.Id, expiration)
	if err != nil {
		return nil, errno.ServerErr(errno.ErrGetUserMsg, err.Error())
	}
	select {
	case <-ctx.Done():
		if err := m.cacher.Rewrite(request.Id, msgs); err != nil {
			return nil, errno.ServerErr(errno.ErrRewriteMsg, err.Error())
		}
	default:
		// TODO: 解决用户连接异常断开后其他用户向其发送离线消息
		// gateway依旧保持连接导致其依旧向chat模块获取消息而写入重复历史消息的问题
		if err := m.cacher.CreateHistory(request.Id, msgs); err != nil {
			return nil, errno.ServerErr(errno.ErrCreateHistory, err.Error())
		}
		response.Msg = msgs
	}
	return &response, nil
}
