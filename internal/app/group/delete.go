package group

import (
	"chat/internal/app/service"
	"chat/internal/pkg/errno"
	"context"
)

func (m *Manager) Delete(ctx context.Context, request *service.DeleteRequest) (*service.Response, error) {
	m.logger.Info("Group service, Delete service")
	group, err := m.localer.GetGroupInfoWithID(request.GroupId)
	if err != nil {
		return nil, errno.ServerErr(errno.ErrDatabase, err.Error())
	}
	if group.UserID != request.UserId {
		return nil, errno.ServerErr(errno.ErrGroupRole, "no role")
	}
	err = m.localer.DeleteGroupWithID(request.GroupId)
	if err != nil {
		return nil, errno.ServerErr(errno.ErrDatabase, err.Error())
	}
	return &service.Response{}, nil
}
