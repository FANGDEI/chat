package group

import (
	"chat/internal/app/service"
	"chat/internal/pkg/errno"
	"context"
)

func (m *Manager) UpdateGroupAvatar(ctx context.Context, request *service.UpdateGroupAvatarRequest) (*service.Response, error) {
	m.logger.Info("Group service, UpdateGroupAvatar service")
	group, err := m.localer.GetGroupInfoWithID(request.GroupId)
	if err != nil {
		return nil, errno.ServerErr(errno.ErrDatabase, err.Error())
	}
	if group.UserID != request.UserId {
		return nil, errno.ServerErr(errno.ErrGroupRole, "no role")
	}
	err = m.localer.UpdateGroupAvatarWithID(request.GroupId, request.Avatar)
	if err != nil {
		return nil, errno.ServerErr(errno.ErrDatabase, err.Error())
	}
	return &service.Response{}, nil
}
