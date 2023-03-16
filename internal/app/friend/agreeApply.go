package friend

import (
	"chat/internal/app/service"
	"chat/internal/pkg/errno"
	"context"
)

func (m *Manager) AgreeApply(ctx context.Context, request *service.AgreeApplyRequest) (*service.Response, error) {
	m.logger.Info("Friend service, AcceptFriend service")
	selfID, friendID, agree := request.Id, request.FriendId, request.Agree
	if agree {
		err := m.localer.CreateFriend(selfID, friendID)
		if err != nil {
			return nil, errno.ServerErr(errno.ErrDatabase, err.Error())
		}
	}
	err := m.localer.DeleteApply(selfID, friendID)
	if err != nil {
		return nil, errno.ServerErr(errno.ErrDatabase, err.Error())
	}
	return &service.Response{}, nil
}
