package chat

import (
	"chat/internal/app/service"
	"chat/internal/pkg/errno"
	"context"
)

func (m *Manager) GetUserHistory(ctx context.Context, request *service.GetUserHistoryRequest) (*service.GetUserHistoryResponse, error) {
	m.logger.Info("Chat service, GetUserHistory service")
	list, err := m.cacher.GetUserHistory(request.UserId, request.OtherId, request.Offset, request.Limit, request.Pagination)
	if err != nil {
		return nil, errno.ServerErr(errno.ErrRedis, err.Error())
	}
	return &service.GetUserHistoryResponse{
		List: list,
	}, nil
}
