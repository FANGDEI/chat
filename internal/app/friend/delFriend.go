package friend

import (
	"chat/internal/app/service"
	"chat/internal/pkg/errno"
	"context"
)

func (m *Manager) DelFriend(ctx context.Context, request *service.DelFriendRequest) (*service.Response, error) {
	m.logger.Info("Friend Service, DelFriend Service")
	id, friendID := request.Id, request.FriendId
	if exists := m.localer.IsFriend(id, friendID); !exists {
		return nil, errno.ServerErr(errno.ErrFriendNotExists, "not friend")
	}
	err := m.localer.DeleteFriend(id, friendID)
	if err != nil {
		return nil, errno.ErrDatabase
	}
	return &service.Response{}, nil
}
