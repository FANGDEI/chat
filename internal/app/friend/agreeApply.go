package friend

import (
	"chat/internal/app/service"
	"chat/internal/pkg/errno"
	"context"
)

func (m *Manager) AgreeApply(ctx context.Context, request *service.AgreeApplyRequest) (*service.Response, error) {
	m.logger.Info("Friend service, AcceptFriend service")
	UserID, friendID, agree := request.UserId, request.FriendId, request.Agree
	if exists := m.localer.ExistApply(UserID, friendID); !exists {
		return nil, errno.ServerErr(errno.ErrFriendApplyNotExists, "apply not exists")
	}
	if agree {
		err := m.localer.CreateFriend(UserID, friendID)
		if err != nil {
			return nil, errno.ServerErr(errno.ErrDatabase, err.Error())
		}
	}
	err := m.localer.DeleteApply(UserID, friendID)
	if err != nil {
		return nil, errno.ServerErr(errno.ErrDatabase, err.Error())
	}
	return &service.Response{}, nil
}
