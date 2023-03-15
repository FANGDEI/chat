package friend

import (
	"chat/internal/app/service"
	"chat/internal/pkg/errno"
	"context"
)

func (m *Manager) AcceptFriend(ctx context.Context, request *service.AcceptFriendRequest) (*service.Response, error) {
	m.logger.Info("Friend service, AcceptFriend service")
	selfUuid, uuid, accept := request.SelfUuid, request.Uuid, request.Agree
	self, err := m.localer.GetUserInfoWithUuid(selfUuid)
	if err != nil {
		return nil, errno.ServerErr(errno.ErrDatabase, err.Error())
	}
	other, err := m.localer.GetUserInfoWithUuid(uuid)
	if err != nil {
		return nil, errno.ServerErr(errno.ErrDatabase, err.Error())
	}
	if accept {
		err := m.localer.AcceptFriend(self.ID, other.ID)
		if err != nil {
			return nil, errno.ServerErr(errno.ErrDatabase, err.Error())
		}
	} else {
		err := m.localer.DeAcceptFriend(self.ID, other.ID)
		if err != nil {
			return nil, errno.ServerErr(errno.ErrDatabase, err.Error())
		}
	}
	return &service.Response{}, nil
}
