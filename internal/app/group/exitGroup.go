package group

import (
	"chat/internal/app/service"
	"chat/internal/pkg/errno"
	"context"
)

func (m *Manager) ExitGroup(ctx context.Context, request *service.ExitGroupRequest) (*service.Response, error) {
	m.logger.Info("Group service, ExitGroup service")
	if exists := m.localer.IsMember(request.UserId, request.GroupId); !exists {
		return nil, errno.ServerErr(errno.ErrGroupMemberNot, "not group member")
	}
	err := m.localer.DeleteGroupMemberWithUserIDAndGroupID(request.UserId, request.GroupId)
	if err != nil {
		return nil, errno.ServerErr(errno.ErrDatabase, err.Error())
	}
	return &service.Response{}, nil
}
