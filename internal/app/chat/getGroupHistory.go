package chat

import (
	"chat/internal/app/service"
	"chat/internal/pkg/errno"
	"context"
)

func (m *Manager) GetGroupHistory(ctx context.Context, request *service.GetGroupHistoryRequest) (*service.GetGroupHistoryResponse, error) {
	m.logger.Info("Chat service, GetGroupHistory service")
	list, err := m.cacher.GetGroupHistory(request.GroupId, request.Offset, request.Limit, request.Pagination)
	if err != nil {
		return nil, errno.ServerErr(errno.ErrRedis, err.Error())
	}
	return &service.GetGroupHistoryResponse{
		List: list,
	}, nil
}
