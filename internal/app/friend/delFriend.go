package friend

import (
	"chat/internal/app/service"
	"chat/internal/pkg/errno"
	"context"
)

func (m *Manager) DelFriend(ctx context.Context, request *service.DelFriendRequest) (*service.Response, error) {
	m.logger.Info("Friend Service, DelFriend Service")
	UserID, friendID := request.UserId, request.FriendId
	if exists := m.localer.IsFriend(UserID, friendID); !exists {
		return nil, errno.ServerErr(errno.ErrFriendNotExists, "not friend")
	}
	err := m.localer.DeleteFriend(UserID, friendID)
	if err != nil {
		return nil, errno.ErrDatabase
	}
	return &service.Response{}, nil
}
