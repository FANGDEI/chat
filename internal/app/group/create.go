package group

import (
	"chat/internal/app/service"
	"chat/internal/pkg/errno"
	"chat/internal/pkg/localer"
	"context"
)

func (m *Manager) Create(ctx context.Context, request *service.CreateRequest) (*service.Response, error) {
	m.logger.Info("Group Service, Create service")
	_, err := m.localer.GetGroupInfoWithName(request.GroupName)
	if err == nil {
		return nil, errno.ServerErr(errno.ErrGroupExists, "group exists")
	}
	err = m.localer.CreateGroup(localer.Group{
		UserID: request.UserId,
		Name:   request.GroupName,
		Avatar: "default.jpg",
	})
	if err != nil {
		return nil, errno.ServerErr(errno.ErrDatabase, err.Error())
	}
	return &service.Response{}, nil
}
